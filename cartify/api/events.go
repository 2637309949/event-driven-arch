package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ThreeDotsLabs/watermill-http/v2/pkg/http"
	watermillHTTP "github.com/ThreeDotsLabs/watermill-http/v2/pkg/http"
	"github.com/ThreeDotsLabs/watermill-redisstream/pkg/redisstream"
	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
	"github.com/ThreeDotsLabs/watermill/message/router/plugin"
	"github.com/redis/go-redis/v9"
)

type PlaceOrderCommand struct {
	TrxId  int64 `json:"trxid"`
	UserId int64 `json:"user_id"`
}

type TrxState struct {
	TrxId    int64  `json:"trxid"`
	Type     int    `json:"type"`
	State    int    `json:"state"`
	Name     string `json:"name"`
	Progress int    `json:"progress"`
}

type TrxStateUpdated struct {
	Type     int    `json:"type"`
	State    string `json:"state"`
	Progress int    `json:"progress"`
}

type Routers struct {
	EventsRouter *message.Router
	CommandBus   *cqrs.CommandBus
	SSERouter    http.SSERouter
}

func (r *Routers) Run(ctx context.Context) {
	ctx, _ = signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	go func() {
		err := r.EventsRouter.Run(ctx)
		if err != nil {
			panic(err)
		}
	}()
	<-r.EventsRouter.Running()

	go func() {
		err := r.SSERouter.Run(ctx)
		if err != nil {
			panic(err)
		}
	}()
	<-r.SSERouter.Running()
	go func() {
		<-ctx.Done()
		if err := r.EventsRouter.Close(); err != nil {
			log.Println("router close error:", err)
		}
	}()
}

func NewRouters(ctx context.Context, cfg *Config, repo *Repository) *Routers {
	redisClient := redis.NewClient(&redis.Options{Addr: redisAddr})
	marshaler := cqrs.JSONMarshaler{
		GenerateName: cqrs.StructName,
	}
	subscriber, err := redisstream.NewSubscriber(redisstream.SubscriberConfig{Client: redisClient}, logger)
	if err != nil {
		panic(err)
	}
	publisher, err := redisstream.NewPublisher(redisstream.PublisherConfig{
		Client: redisClient,
	}, logger)
	if err != nil {
		panic(err)
	}
	router, err := message.NewRouter(message.RouterConfig{}, logger)
	if err != nil {
		panic(err)
	}
	router.AddPlugin(plugin.SignalsHandler)
	router.AddMiddleware(middleware.Recoverer)
	eventProcessor, err := cqrs.NewEventProcessorWithConfig(router, cqrs.EventProcessorConfig{
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
	})
	if err != nil {
		panic(err)
	}
	err = eventProcessor.AddHandlers(
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
	)
	if err != nil {
		panic(err)
	}
	commandBus, err := cqrs.NewCommandBusWithConfig(publisher, cqrs.CommandBusConfig{
		GeneratePublishTopic: func(params cqrs.CommandBusGeneratePublishTopicParams) (string, error) {
			return params.CommandName, nil
		},
		Marshaler: marshaler,
		Logger:    logger,
	})
	if err != nil {
		panic(err)
	}

	sseRouter, err := watermillHTTP.NewSSERouter(
		watermillHTTP.SSERouterConfig{
			UpstreamSubscriber: subscriber,
			ErrorHandler:       watermillHTTP.DefaultErrorHandler,
		},
		logger,
	)
	if err != nil {
		panic(err)
	}

	routers := Routers{}
	routers.EventsRouter = router
	routers.CommandBus = commandBus
	routers.SSERouter = sseRouter
	return &routers
}
