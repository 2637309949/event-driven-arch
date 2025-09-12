package main

import (
	"context"
	"net/http"

	"github.com/ThreeDotsLabs/watermill"
	watermillHTTP "github.com/ThreeDotsLabs/watermill-http/v2/pkg/http"
	"github.com/ThreeDotsLabs/watermill-redisstream/pkg/redisstream"
	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

var (
	topic       = "OrderPlaced"
	ctx         = context.Background()
	logger      = watermill.NewStdLogger(false, false)
	redisClient = redis.NewClient(&redis.Options{Addr: "127.0.0.1:6379"})
	marshaler   = cqrs.JSONMarshaler{
		GenerateName: cqrs.StructName,
	}
)

func main() {
	subscriber := mustNew(redisstream.NewSubscriber(
		redisstream.SubscriberConfig{
			Client: redisClient,
		},
		logger,
	)).(*redisstream.Subscriber)
	publisher := mustNew(redisstream.NewPublisher(redisstream.PublisherConfig{
		Client: redisClient,
	}, logger)).(*redisstream.Publisher)
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
	orderplacedStream := orderplacedStreamAdapter{}
	messagesHandler := sseRouter.AddHandler(topic, orderplacedStream)
	mustRoutine(func() (err error) {
		err = sseRouter.Run(ctx)
		return
	})

	r := gin.Default()
	r.Use(requestIDMiddleware())
	r.GET("/api/v1/orderplaced", wrapHttpHandler(messagesHandler))
	r.POST("/api/v1/order", func(c *gin.Context) {
		var placeOrderCommand PlaceOrderCommand
		if err := c.ShouldBindJSON(&placeOrderCommand); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": err.Error(),
			})
			return
		}
		placeOrderCommand = newFakePlaceOrderCommand(placeOrderCommand.RequestID)
		mustCall(commandBus.Send(ctx, placeOrderCommand))
		c.JSON(http.StatusOK, gin.H{
			"status":     "ok",
			"request_id": placeOrderCommand.RequestID,
		})
	})
	r.Run()
}
