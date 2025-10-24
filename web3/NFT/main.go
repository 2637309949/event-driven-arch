package main

import (
	"context"

	"github.com/ThreeDotsLabs/watermill"
	_ "github.com/lib/pq"
	"github.com/robfig/cron"
)

var (
	topic          = "EventOccurred"
	redisAddr      = "127.0.0.1:6379"
	driverName     = "postgres"
	dataSourceName = "postgres://:@127.0.0.1:5432/web3?sslmode=disable"
	ctx            = context.Background()
	logger         = watermill.NewStdLogger(false, false)
	db             = Open(driverName, dataSourceName)
)

func main() {
	MigrateDB(db)
	config := NewConfig()
	repos := NewRepository(db)
	routers := NewRouters(ctx, config, repos)
	ct := NewContract(routers)
	ct.Watch(ctx)

	c := cron.New()
	t := NewTimed(c, ct, repos, routers.EventBus)

	t.Start(ctx)
	srv := NewHandler(repos, routers)
	srv.Run(ctx)
}
