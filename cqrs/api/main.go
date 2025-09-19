package main

import (
	"context"
	stdSQL "database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	watermillHTTP "github.com/ThreeDotsLabs/watermill-http/v2/pkg/http"
	"github.com/ThreeDotsLabs/watermill-redisstream/pkg/redisstream"
	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
	"github.com/ThreeDotsLabs/watermill/message/router/plugin"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
)

var (
	topic       = "TrxState"
	ctx         = context.Background()
	logger      = watermill.NewStdLogger(false, false)
	redisClient = redis.NewClient(&redis.Options{Addr: "127.0.0.1:6379"})
	db          = mustNew(stdSQL.Open("postgres", "postgres://Doubl:@127.0.0.1:5432/testdb?sslmode=disable")).(*stdSQL.DB)
	epoch       = time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC).UnixNano() / int64(time.Millisecond) // 例：epoch 设为 2020-01-01 00:00:00 UTC 的毫秒数
	sf          = mustNew(NewSnowflake(1, epoch)).(*Snowflake)                                     // nodeID = 1
	marshaler   = cqrs.JSONMarshaler{
		GenerateName: cqrs.StructName,
	}
)

func main() {
	mustCall(MigrateDB(db))
	repo := NewRepository(db)

	subscriber := mustNew(redisstream.NewSubscriber(redisstream.SubscriberConfig{Client: redisClient}, logger)).(*redisstream.Subscriber)
	publisher := mustNew(redisstream.NewPublisher(redisstream.PublisherConfig{
		Client: redisClient,
	}, logger)).(*redisstream.Publisher)
	router := mustNew(message.NewRouter(message.RouterConfig{}, logger)).(*message.Router)
	router.AddPlugin(plugin.SignalsHandler)
	router.AddMiddleware(middleware.Recoverer)
	eventProcessor := mustNew(cqrs.NewEventProcessorWithConfig(router, cqrs.EventProcessorConfig{
		GenerateSubscribeTopic: func(params cqrs.EventProcessorGenerateSubscribeTopicParams) (string, error) {
			return params.EventName, nil
		},
		SubscriberConstructor: func(params cqrs.EventProcessorSubscriberConstructorParams) (message.Subscriber, error) {
			return redisstream.NewSubscriber(redisstream.SubscriberConfig{
				Client:        redisClient,
				ConsumerGroup: params.HandlerName,
			}, logger)
		},
		Marshaler: marshaler,
		Logger:    logger,
	})).(*cqrs.EventProcessor)
	mustCall(eventProcessor.AddHandlers(
		cqrs.NewEventHandler(
			"OnOrderPlacedHandler",
			func(ctx context.Context, ev *TrxState) error {
				fmt.Printf("received event %+v\n", ev)
				err := repo.SaveTrx(ctx, ev.TrxId, func(trx *Trx) {
					trx.Type = ev.Type
					trx.State = ev.State
					trx.Name = ev.Name
					trx.Progress = ev.Progress
				})
				return err
			},
		),
	))
	commandBus := mustNew(cqrs.NewCommandBusWithConfig(publisher, cqrs.CommandBusConfig{
		GeneratePublishTopic: func(params cqrs.CommandBusGeneratePublishTopicParams) (string, error) {
			return params.CommandName, nil
		},
		Marshaler: marshaler,
		Logger:    logger,
	})).(*cqrs.CommandBus)
	sseRouter := mustNew(watermillHTTP.NewSSERouter(
		watermillHTTP.SSERouterConfig{
			UpstreamSubscriber: subscriber,
			ErrorHandler:       watermillHTTP.DefaultErrorHandler,
		},
		logger,
	)).(watermillHTTP.SSERouter)
	trxStream := trxStreamAdapter{logger: logger, repo: repo}
	trxHandler := sseRouter.AddHandler(topic, trxStream)
	mustRoutine(func() (err error) {
		err = sseRouter.Run(ctx)
		return
	})
	mustRoutine(func() (err error) {
		err = router.Run(ctx)
		return
	})

	r := gin.Default()
	r.Use(requestIDMiddleware())
	r.GET("/trx/:id", func(c *gin.Context) {
		c.Request.SetPathValue("id", c.Param("id"))
		trxHandler(c.Writer, c.Request)
	})
	r.POST("/orders", func(c *gin.Context) {
		var placeOrderCommand PlaceOrderCommand
		if err := c.ShouldBindJSON(&placeOrderCommand); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": err.Error(),
			})
			return
		}
		placeOrderCommand = newFakePlaceOrderCommand(placeOrderCommand.UserId)
		placeOrderCommand.TrxId = sf.NextID()
		mustCall(commandBus.Send(ctx, placeOrderCommand))
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"trxid":  placeOrderCommand.TrxId,
		})
	})
	r.Run()
}
