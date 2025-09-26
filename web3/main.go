package main

import (
	"context"
	stdSQL "database/sql"
	"fmt"
	"net/http"

	"github.com/ThreeDotsLabs/watermill"
	_ "github.com/lib/pq"
)

type Config struct {
}

var (
	topic  = "ItemSet"
	ctx    = context.Background()
	logger = watermill.NewStdLogger(false, false)
)

func main() {
	config := Config{}
	db, err := stdSQL.Open("postgres", "postgres://Doubl:@127.0.0.1:5432/testdb?sslmode=disable")
	if err != nil {
		panic(err)
	}
	err = MigrateDB(db)
	if err != nil {
		fmt.Println(err)
	}
	repos := NewRepository(db)

	routers, err := NewRouters(ctx, &config, repos)
	if err != nil {
		panic(err)
	}
	ct := NewContract(routers.EventBus, routers.CommandBus)
	err = ct.Watch(ctx)
	if err != nil {
		panic(err)
	}

	mux := NewHandler(repos, routers.EventBus, routers.SSERouter)
	server := &http.Server{Addr: ":8080", Handler: mux}
	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
