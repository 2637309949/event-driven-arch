package main

import (
	"context"

	_ "github.com/lib/pq"

	"github.com/ThreeDotsLabs/watermill"
)

var (
	redisAddr      = "127.0.0.1:6379"
	driverName     = "postgres"
	dataSourceName = "postgres://:@127.0.0.1:5432/testdb?sslmode=disable"
	ctx            = context.Background()
	logger         = watermill.NewStdLogger(false, false)
	db             = Open(driverName, dataSourceName)
)

func main() {
	MigrateDB(db)
	config := NewConfig()
	repo := NewRepository(db)
	routers := NewRouters(ctx, config, repo)
	routers.Run(ctx)
}
