package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"web3/contract/store"

	watermillHTTP "github.com/ThreeDotsLabs/watermill-http/v2/pkg/http"
	"github.com/gin-gonic/gin"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-redisstream/pkg/redisstream"
	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
	"github.com/ThreeDotsLabs/watermill/message/router/plugin"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"
	"github.com/redis/go-redis/v9"
)

var (
	topic           = "ItemSet"
	ctx             = context.Background()
	logger          = watermill.NewStdLogger(false, false)
	redisClient     = redis.NewClient(&redis.Options{Addr: "127.0.0.1:6379"})
	marshaler       = cqrs.JSONMarshaler{GenerateName: cqrs.StructName}
	client          = mustNew(ethclient.Dial("wss://sepolia.infura.io/ws/v3/b0daa49c16d7466cbdf68176ba2a243a")).(*ethclient.Client)
	contractAddress = common.HexToAddress("0xfeadcf82070998D19A215C91E19638Bfcd1Ab854")
	storeInstance   = mustNew(store.NewStore(contractAddress, client)).(*store.Store)
	evc             = make(chan *store.StoreItemSet)
	sub             = mustNew(storeInstance.WatchItemSet(&bind.WatchOpts{Context: ctx}, evc)).(event.Subscription)
	subscriber      = mustNew(redisstream.NewSubscriber(redisstream.SubscriberConfig{Client: redisClient}, logger)).(*redisstream.Subscriber)
	publisher       = mustNew(redisstream.NewPublisher(redisstream.PublisherConfig{Client: redisClient}, logger)).(*redisstream.Publisher)
	router          = mustNew(message.NewRouter(message.RouterConfig{}, logger)).(*message.Router)
	eventBus        = mustNew(cqrs.NewEventBusWithConfig(publisher, cqrs.EventBusConfig{
		GeneratePublishTopic: func(params cqrs.GenerateEventPublishTopicParams) (string, error) {
			return params.EventName, nil
		},
		Marshaler: marshaler,
		Logger:    logger,
	})).(*cqrs.EventBus)
	eventProcessor = mustNew(cqrs.NewEventProcessorWithConfig(router, cqrs.EventProcessorConfig{
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
	commandProcessor = mustNew(cqrs.NewCommandProcessorWithConfig(router, cqrs.CommandProcessorConfig{
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
	sseRouter = mustNew(watermillHTTP.NewSSERouter(
		watermillHTTP.SSERouterConfig{
			UpstreamSubscriber: subscriber,
			ErrorHandler:       watermillHTTP.DefaultErrorHandler,
		},
		logger,
	)).(watermillHTTP.SSERouter)
	itemStream  = itemStreamAdapter{logger: logger}
	itemHandler = sseRouter.AddHandler(topic, itemStream)
)

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	router.AddPlugin(plugin.SignalsHandler)
	router.AddMiddleware(middleware.Recoverer)
	mustCall(commandProcessor.AddHandlers(
		cqrs.NewCommandHandler(
			"ItemSetCommandHandler",
			func(ctx context.Context, cmd *ItemSetCommand) error {
				is := ItemSet{}
				is.Key = strings.TrimRight(string(cmd.Key[:]), "\x00")
				is.Value = strings.TrimRight(string(cmd.Value[:]), "\x00")
				return eventBus.Publish(ctx, &is)
			},
		),
	))
	mustCall(eventProcessor.AddHandlers(
		cqrs.NewEventHandler(
			"ItemSetHandler",
			func(ctx context.Context, ev *ItemSet) error {
				fmt.Printf("ItemSet key=%v, value=%v\n", ev.Key, ev.Value)
				return nil
			},
		),
	))
	mustRoutine(func() (err error) {
		err = sseRouter.Run(ctx)
		return
	})
	mustRoutine(func() (err error) {
		err = router.Run(ctx)
		return
	})

	r := gin.Default()
	r.GET("/subscribe/event", func(c *gin.Context) {
		itemHandler(c.Writer, c.Request)
	})
	mustRoutine(func() (err error) {
		err = r.Run()
		return
	})

	for {
		select {
		case ev := <-evc:
			cmd := ItemSetCommand{}
			cmd.Key = ev.Key
			cmd.Value = ev.Value
			eventBus.Publish(ctx, cmd)
		case <-sigs:
			subscriber.Close()
			publisher.Close()
			fmt.Println("订阅退出...")
		case err := <-sub.Err():
			log.Fatal("sub err:", err)
		}
	}
}
