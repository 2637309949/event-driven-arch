package main

import (
	"context"

	"github.com/ThreeDotsLabs/watermill"
	_ "github.com/lib/pq"
)

type Config struct {
}

var (
	topic          = "TrxState"
	redisAddr      = "127.0.0.1:6379"
	driverName     = "postgres"
	dataSourceName = "postgres://:@127.0.0.1:5432/testdb?sslmode=disable"
	ctx            = context.Background()
	logger         = watermill.NewStdLogger(false, false)
	db             = Open(driverName, dataSourceName)
)

func main() {
	MigrateDB(db)
	cfg := Config{}
	repo := NewRepository(db)
	routers := NewRouters(ctx, &cfg, repo)
	srv := NewHandler(repo, routers)
	srv.Run(ctx)
}
