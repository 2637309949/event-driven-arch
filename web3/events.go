package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
	"web3/contract/nft"

	"github.com/ThreeDotsLabs/watermill-http/v2/pkg/http"
	"github.com/ThreeDotsLabs/watermill-redisstream/pkg/redisstream"
	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
	"github.com/ThreeDotsLabs/watermill/message/router/plugin"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/redis/go-redis/v9"
)

type ItemSetCommand struct {
	Key   [32]byte
	Value [32]byte
}

type ItemSet struct {
	Key   string
	Value string
}

type Routers struct {
	EventsRouter *message.Router
	EventBus     *cqrs.EventBus
	CommandBus   *cqrs.CommandBus
	SSERouter    http.SSERouter
}

type EventLog struct {
	TxHash          string
	LogIndex        int
	BlockNumber     int64
	BlockTime       time.Time
	ContractAddress string
	EventSignature  string
	EventName       string
	Topic0          string
	Topic1          string
	Topic2          string
	Topic3          string
	Data            string
	CreatedAt       time.Time
}

func NewRouters(ctx context.Context, cfg *Config, repo *Repository) (*Routers, error) {
	ctx, _ = signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	redisClient := redis.NewClient(&redis.Options{Addr: "127.0.0.1:6379"})
	marshaler := cqrs.JSONMarshaler{GenerateName: cqrs.StructName}
	routers := Routers{}
	router, err := message.NewRouter(message.RouterConfig{}, logger)
	if err != nil {
		return nil, err
	}
	router.AddPlugin(plugin.SignalsHandler)
	router.AddMiddleware(middleware.Recoverer)
	subscriber, err := redisstream.NewSubscriber(redisstream.SubscriberConfig{Client: redisClient}, logger)
	if err != nil {
		return nil, err
	}
	publisher, err := redisstream.NewPublisher(redisstream.PublisherConfig{Client: redisClient}, logger)
	if err != nil {
		return nil, err
	}
	eventBus, err := cqrs.NewEventBusWithConfig(publisher, cqrs.EventBusConfig{
		GeneratePublishTopic: func(params cqrs.GenerateEventPublishTopicParams) (string, error) {
			return params.EventName, nil
		},
		Marshaler: marshaler,
		Logger:    logger,
	})
	if err != nil {
		return nil, err
	}
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
		return nil, err
	}
	eventProcessor.AddHandlers(
		cqrs.NewEventHandler(
			"ItemSetHandler",
			func(ctx context.Context, ev *ItemSet) error {
				fmt.Printf("ItemSet key=%v, value=%v\n", ev.Key, ev.Value)
				return nil
			},
		),
		cqrs.NewEventHandler(
			"EventParsedHandler",
			func(ctx context.Context, ev *EventLog) error {
				fmt.Printf("EventParsed EventName=%v, TxHash=%v\n", ev.EventName, ev.TxHash)
				eventName := ev.EventName
				switch eventName {
				case "Transfer":
					contractAbi, _ := abi.JSON(strings.NewReader(nft.NftABI))
					from := common.HexToAddress(ev.Topic1)
					to := common.HexToAddress(ev.Topic2)
					var out []interface{}
					err := contractAbi.UnpackIntoInterface(&out, "Transfer", []byte(ev.Data))
					if err != nil {
						log.Println("Unpack err:", err)
						return err
					}
					var valueOrTokenId *big.Int
					if len(out) > 0 {
						valueOrTokenId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
					}
					ep := EventParsed{
						TxHash:          ev.TxHash,
						LogIndex:        ev.LogIndex,
						BlockNumber:     ev.BlockNumber,
						BlockTime:       ev.BlockTime,
						ContractAddress: ev.ContractAddress,
						EventName:       ev.EventName,
						FromAddress:     from.Hex(),
						ToAddress:       to.Hex(),
						TokenID:         valueOrTokenId,
						Value:           valueOrTokenId,
					}
					err = repo.InsertEventParsed(ctx, &ep)
					if err != nil {
						log.Println("InsertEventParsed err:", err)
						return err
					}
				}
				return nil
			},
		),
		cqrs.NewEventHandler(
			"EventRawHandler",
			func(ctx context.Context, ev *EventLog) error {
				fmt.Printf("EventRaw EventName=%v, TxHash=%v\n", ev.EventName, ev.TxHash)
				er := EventRaw{
					TxHash:          ev.TxHash,
					LogIndex:        ev.LogIndex,
					BlockNumber:     ev.BlockNumber,
					BlockTime:       ev.BlockTime,
					ContractAddress: ev.ContractAddress,
					EventSignature:  ev.EventSignature,
					EventName:       ev.EventName,
					Topic0:          ev.Topic0,
					Topic1:          ev.Topic1,
					Topic2:          ev.Topic2,
					Topic3:          ev.Topic3,
					Data:            ev.Data,
					CreatedAt:       ev.CreatedAt,
				}
				err := repo.InsertEventRaw(ctx, &er)
				if err != nil {
					log.Println("InsertEventRaw err:", err)
					return err
				}
				return nil
			},
		),
	)
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
	commandProcessor, err := cqrs.NewCommandProcessorWithConfig(router, cqrs.CommandProcessorConfig{
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
	})
	if err != nil {
		return nil, err
	}
	commandProcessor.AddHandlers(
		cqrs.NewCommandHandler(
			"ItemSetCommandHandler",
			func(ctx context.Context, cmd *ItemSetCommand) error {
				is := ItemSet{}
				is.Key = strings.TrimRight(string(cmd.Key[:]), "\x00")
				is.Value = strings.TrimRight(string(cmd.Value[:]), "\x00")
				return eventBus.Publish(ctx, &is)
			},
		),
	)
	sseRouter, err := http.NewSSERouter(
		http.SSERouterConfig{
			UpstreamSubscriber: subscriber,
			ErrorHandler:       http.DefaultErrorHandler,
		},
		logger,
	)
	if err != nil {
		return nil, err
	}

	go func() {
		err := router.Run(ctx)
		if err != nil {
			panic(err)
		}
	}()
	<-router.Running()

	go func() {
		err := sseRouter.Run(ctx)
		if err != nil {
			panic(err)
		}
	}()
	<-sseRouter.Running()

	go func() {
		<-ctx.Done()
		if err := router.Close(); err != nil {
			log.Println("router close error:", err)
		}
	}()

	routers.SSERouter = sseRouter
	routers.EventsRouter = router
	routers.EventBus = eventBus
	routers.CommandBus = commandBus
	return &routers, nil
}
