package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ThreeDotsLabs/watermill-redisstream/pkg/redisstream"
	"github.com/ThreeDotsLabs/watermill-sql/v4/pkg/sql"
	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
	"github.com/redis/go-redis/v9"
)

type UploadFileCommand struct {
	TrxId    int64  `json:"trxid"`
	SavePath string `json:"save_path"`
	NewName  string `json:"new_name"`
	OrigName string `json:"orig_name"`
	Ext      string `json:"ext"`
	MimeType string `json:"mime_type"`
}

type FileParsed struct {
	FileId   int64  `json:"file_id"`
	TrxId    int64  `json:"trxid"`
	FilePath string `json:"file_path"`
	FileName string `json:"file_name"`
	Ext      string `json:"ext"`
	MimeType string `json:"mime_type"`
}

type TrxState struct {
	TrxId    int64  `json:"trxid"`
	Type     int    `json:"type"`
	State    int    `json:"state"`
	Name     string `json:"name"`
	Progress int    `json:"progress"`
}

type Routers struct {
	EventsRouter    *message.Router
	EventBus        *cqrs.EventBus
	delayedRequeuer *sql.DelayedRequeuer
}

func (r *Routers) Run(ctx context.Context) error {
	ctx, _ = signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	go func() {
		err := r.EventsRouter.Run(ctx)
		if err != nil {
			panic(err)
		}
	}()
	<-r.EventsRouter.Running()

	go func() {
		err := r.delayedRequeuer.Run(ctx)
		if err != nil {
			panic(err)
		}
	}()

	<-ctx.Done()
	if err := r.EventsRouter.Close(); err != nil {
		log.Println("router close error:", err)
		return err
	}
	return nil
}

func NewRouters(ctx context.Context, cfg *Config, repo *Repository) (*Routers, error) {
	marshaler := cqrs.JSONMarshaler{
		GenerateName: cqrs.StructName,
	}
	redisClient := redis.NewClient(&redis.Options{Addr: redisAddr})
	publisher, err := redisstream.NewPublisher(redisstream.PublisherConfig{
		Client: redisClient,
	}, logger)
	if err != nil {
		panic(err)
	}
	delayedRequeuer, err := sql.NewPostgreSQLDelayedRequeuer(sql.DelayedRequeuerConfig{
		DB:        sql.BeginnerFromStdSQL(repo.db),
		Publisher: publisher,
		DelayOnError: &middleware.DelayOnError{
			InitialInterval: 10 * time.Second,
			MaxInterval:     3 * time.Minute,
			Multiplier:      2,
		},
		Logger: logger,
	})
	if err != nil {
		panic(err)
	}

	eventBus, err := cqrs.NewEventBusWithConfig(publisher, cqrs.EventBusConfig{
		GeneratePublishTopic: func(params cqrs.GenerateEventPublishTopicParams) (string, error) {
			return params.EventName, nil
		},
		Marshaler: marshaler,
		Logger:    logger,
	})
	if err != nil {
		panic(err)
	}
	router := message.NewDefaultRouter(logger)
	router.AddMiddleware(delayedRequeuer.Middleware()...)
	// router.AddMiddleware(notifyMiddleware(publisher))
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
			"OrderPlacedHandler",
			func(ctx context.Context, ev *FileParsed) error {
				err := repo.CreateFile(ctx, &File{FileId: ev.FileId, State: 101, UserId: ev.UserId})
				if err != nil {
					return err
				}
				trxState := TrxState{}
				trxState.TrxId = ev.TrxId
				trxState.Type = 102
				trxState.State = 1006
				trxState.Name = "待支付"
				trxState.Progress = 100
				eventBus.Publish(ctx, &trxState)
				time.Sleep(3 * time.Second)
				return err
			},
		),
	)
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
		panic(err)
	}
	err = commandProcessor.AddHandlers(
		cqrs.NewCommandHandler(
			"PlaceOrderHandler",
			func(ctx context.Context, cmd *UploadFileCommand) error {
				// 接受请求
				trxState := TrxState{}
				trxState.TrxId = cmd.TrxId
				trxState.Type = 102
				trxState.State = 1001
				trxState.Name = "接受请求"
				trxState.Progress = 10
				eventBus.Publish(ctx, &trxState)
				time.Sleep(3 * time.Second)
				// 正在锁定库存
				trxState = TrxState{}
				trxState.TrxId = cmd.TrxId
				trxState.Type = 102
				trxState.State = 1002
				trxState.Name = "正在锁定库存"
				trxState.Progress = 20
				eventBus.Publish(ctx, &trxState)
				time.Sleep(3 * time.Second)
				// 锁定库存成功
				trxState = TrxState{}
				trxState.TrxId = cmd.TrxId
				trxState.Type = 102
				trxState.State = 1003
				trxState.Name = "锁定库存成功"
				trxState.Progress = 30
				eventBus.Publish(ctx, &trxState)
				time.Sleep(3 * time.Second)
				// 计算订单价格
				trxState = TrxState{}
				trxState.TrxId = cmd.TrxId
				trxState.Type = 102
				trxState.State = 1004
				trxState.Name = "正在计算订单价格"
				trxState.Progress = 40
				eventBus.Publish(ctx, &trxState)
				time.Sleep(3 * time.Second)
				// 生成待支付单
				trxState = TrxState{}
				trxState.TrxId = cmd.TrxId
				trxState.Type = 102
				trxState.State = 1005
				trxState.Name = "生成待支付单"
				trxState.Progress = 50
				eventBus.Publish(ctx, &trxState)
				time.Sleep(3 * time.Second)
				fileParsed := FileParsed{
					TrxId:   cmd.TrxId,
					OrderId: NextID(),
					UserId:  cmd.UserId,
				}
				return eventBus.Publish(ctx, &fileParsed)
			},
		),
	)
	if err != nil {
		panic(err)
	}

	routers := Routers{}
	routers.delayedRequeuer = delayedRequeuer
	routers.EventsRouter = router
	routers.EventBus = eventBus
	return &routers, nil
}
