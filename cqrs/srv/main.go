package main

import (
	"context"
	stdSQL "database/sql"
	"fmt"
	"time"

	"github.com/ThreeDotsLabs/watermill/message/router/middleware"

	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-redisstream/pkg/redisstream"
	"github.com/ThreeDotsLabs/watermill-sql/v4/pkg/sql"
	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/ThreeDotsLabs/watermill/message"
)

var (
	topic       = "OrderPlaced"
	ctx         = context.Background()
	logger      = watermill.NewStdLogger(false, false)
	redisClient = redis.NewClient(&redis.Options{Addr: "127.0.0.1:6379"})
	db          = mustNew(stdSQL.Open("postgres", "postgres://Doubl:@127.0.0.1:5432/testdb?sslmode=disable")).(*stdSQL.DB)
	marshaler   = cqrs.JSONMarshaler{
		GenerateName: cqrs.StructName,
	}
)

func main() {
	publisher := mustNew(redisstream.NewPublisher(redisstream.PublisherConfig{
		Client: redisClient,
	}, logger)).(*redisstream.Publisher)

	delayedRequeuer := mustNew(sql.NewPostgreSQLDelayedRequeuer(sql.DelayedRequeuerConfig{
		DB:        sql.BeginnerFromStdSQL(db),
		Publisher: publisher,
		DelayOnError: &middleware.DelayOnError{
			InitialInterval: 10 * time.Second,
			MaxInterval:     3 * time.Minute,
			Multiplier:      2,
		},
		Logger: logger,
	})).(*sql.DelayedRequeuer)

	eventBus := mustNew(cqrs.NewEventBusWithConfig(publisher, cqrs.EventBusConfig{
		GeneratePublishTopic: func(params cqrs.GenerateEventPublishTopicParams) (string, error) {
			return params.EventName, nil
		},
		Marshaler: marshaler,
		Logger:    logger,
	})).(*cqrs.EventBus)

	router := message.NewDefaultRouter(logger)
	router.AddMiddleware(delayedRequeuer.Middleware()...)
	router.AddMiddleware(notifyMiddleware(publisher))
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

	commandProcessor := mustNew(cqrs.NewCommandProcessorWithConfig(router, cqrs.CommandProcessorConfig{
		GenerateSubscribeTopic: func(params cqrs.CommandProcessorGenerateSubscribeTopicParams) (string, error) {
			return params.CommandName, nil
		},
		SubscriberConstructor: func(params cqrs.CommandProcessorSubscriberConstructorParams) (message.Subscriber, error) {
			return redisstream.NewSubscriber(redisstream.SubscriberConfig{
				Client:        redisClient,
				ConsumerGroup: params.HandlerName,
			}, logger)
		},
		Marshaler: marshaler,
		Logger:    logger,
	})).(*cqrs.CommandProcessor)

	mustCall(eventProcessor.AddHandlers(
		cqrs.NewEventHandler(
			"OnOrderPlacedHandler",
			func(ctx context.Context, event *OrderPlaced) error {
				if event.OrderID == "" {
					fmt.Println("ERROR: Received order placed without order_id")
					return fmt.Errorf("empty order_id")
				}
				fmt.Printf("Received %v order placed\n", event.Customer.Name)
				return nil
			},
		),
	))

	i := 0
	mustCall(commandProcessor.AddHandlers(
		cqrs.NewCommandHandler(
			"PlaceOrderHandler",
			func(ctx context.Context, cmd *PlaceOrderCommand) error {
				// 处理写操作：比如存数据库
				fmt.Printf("Handling PlaceOrderCommand for customer: %v\n", cmd.Customer.Name)
				// 处理完成后触发 OrderPlaced 事件
				e := OrderPlaced{
					RequestID: cmd.RequestID,
					OrderID:   watermill.NewUUID(),
					Customer:  cmd.Customer,
					Products:  cmd.Products,
					Address:   cmd.Address,
				}
				i++
				if i == 10 {
					e.OrderID = ""
					i = 0
				}
				return eventBus.Publish(ctx, &e)
			},
		),
	))

	mustRoutine(func() (err error) {
		err = delayedRequeuer.Run(ctx)
		return
	})
	mustCall(router.Run(ctx))
}
