package main

import (
	"context"
	stdSQL "database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/ThreeDotsLabs/watermill"
)

type Config struct {
}

var (
	ctx       = context.Background()
	logger    = watermill.NewStdLogger(false, false)
	redisAddr = "127.0.0.1:6379"
)

func main() {
	config := Config{}
	db, err := stdSQL.Open("postgres", "postgres://:@127.0.0.1:5432/testdb?sslmode=disable")
	if err != nil {
		panic(err)
	}

	err = MigrateDB(db)
	if err != nil {
		fmt.Println(err)
	}
	repo := NewRepository(db)
	routers, err := NewRouters(ctx, &config, repo)
	if err != nil {
		panic(err)
	}

	err = routers.Run(ctx)
	if err != nil {
		panic(err)
	}
}
