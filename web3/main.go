package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	_ "github.com/lib/pq"
	"github.com/robfig/cron"
)

type Config struct {
}

var (
	topic  = "EventOccurred"
	ctx    = context.Background()
	logger = watermill.NewStdLogger(false, false)
	redisAddr = "127.0.0.1:6379"
	dbUri  = "postgres://:@127.0.0.1:5432/web3?sslmode=disable"
)

func main() {
	c := cron.New()
	config := Config{}
	db, err := sql.Open("postgres", dbUri)
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

	ct := NewContract(routers)
	err = ct.Watch(ctx)
	if err != nil {
		panic(err)
	}

	t := NewTimed(c, ct, repos, routers.EventBus)
	t.Start(ctx)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	mux := NewHandler(repos, routers.EventBus, routers.SSERouter)
	server := &http.Server{Addr: ":8080", Handler: mux}
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("ListenAndServe error", err, watermill.LogFields{})
		}
	}()
	routers.Run(ctx) // 确保注册完事件处理函数
	logger.Info("Server started at", watermill.LogFields{"port": "8080"})
	<-quit
	logger.Info("Shutting down server...", watermill.LogFields{})
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		logger.Error("ListenAndServe error", err, watermill.LogFields{})
	}
	logger.Info("Server exiting", watermill.LogFields{})
}
