package main

import (
	"context"
	stdSQL "database/sql"
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
	ctx         = context.Background()
	logger      = watermill.NewStdLogger(false, false)
	redisClient = redis.NewClient(&redis.Options{Addr: "127.0.0.1:6379"})
	db          = mustNew(stdSQL.Open("postgres", "postgres://Doubl:@127.0.0.1:5432/testdb?sslmode=disable")).(*stdSQL.DB)
	epoch       = time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC).UnixNano() / int64(time.Millisecond) // 例：epoch 设为 2020-01-01 00:00:00 UTC 的毫秒数
	sf          = mustNew(NewSnowflake(2, epoch)).(*Snowflake)
	marshaler   = cqrs.JSONMarshaler{
		GenerateName: cqrs.StructName,
	}
)

func main() {
	mustCall(MigrateDB(db))
	repo := NewRepository(db)

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
	// router.AddMiddleware(notifyMiddleware(publisher))
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
			func(ctx context.Context, ev *OrderPlaced) error {
				err := repo.CreateOrder(ctx, &Order{OrderId: ev.OrderId, State: 101, UserId: ev.UserId})
				if err != nil {
					return err
				}
				trxState := TrxState{}
				trxState.TrxId = ev.TrxId
				trxState.Type = 101
				trxState.State = 1006
				trxState.Name = "待支付"
				trxState.Progress = 100
				eventBus.Publish(ctx, &trxState)
				time.Sleep(3 * time.Second)
				return err
			},
		),
	))

	mustCall(commandProcessor.AddHandlers(
		cqrs.NewCommandHandler(
			"PlaceOrderHandler",
			func(ctx context.Context, cmd *PlaceOrderCommand) error {
				// 接受请求
				trxState := TrxState{}
				trxState.TrxId = cmd.TrxId
				trxState.Type = 101
				trxState.State = 1001
				trxState.Name = "接受请求"
				trxState.Progress = 10
				eventBus.Publish(ctx, &trxState)
				time.Sleep(3 * time.Second)
				// 正在锁定库存
				trxState = TrxState{}
				trxState.TrxId = cmd.TrxId
				trxState.Type = 101
				trxState.State = 1002
				trxState.Name = "正在锁定库存"
				trxState.Progress = 20
				eventBus.Publish(ctx, &trxState)
				time.Sleep(3 * time.Second)
				// 锁定库存成功
				trxState = TrxState{}
				trxState.TrxId = cmd.TrxId
				trxState.Type = 101
				trxState.State = 1003
				trxState.Name = "锁定库存成功"
				trxState.Progress = 30
				eventBus.Publish(ctx, &trxState)
				time.Sleep(3 * time.Second)
				// 计算订单价格
				trxState = TrxState{}
				trxState.TrxId = cmd.TrxId
				trxState.Type = 101
				trxState.State = 1004
				trxState.Name = "正在计算订单价格"
				trxState.Progress = 40
				eventBus.Publish(ctx, &trxState)
				time.Sleep(3 * time.Second)
				// 生成待支付单
				trxState = TrxState{}
				trxState.TrxId = cmd.TrxId
				trxState.Type = 101
				trxState.State = 1005
				trxState.Name = "生成待支付单"
				trxState.Progress = 50
				eventBus.Publish(ctx, &trxState)
				time.Sleep(3 * time.Second)
				orderPlaced := OrderPlaced{
					TrxId:   cmd.TrxId,
					OrderId: sf.NextID(),
					UserId:  cmd.UserId,
				}
				return eventBus.Publish(ctx, &orderPlaced)
			},
		),
	))

	mustRoutine(func() (err error) {
		err = delayedRequeuer.Run(ctx)
		return
	})
	mustCall(router.Run(ctx))
}
