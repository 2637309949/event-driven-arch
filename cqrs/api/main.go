package main

import (
	"context"
	stdSQL "database/sql"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	_ "github.com/lib/pq"
)

type Config struct {
}

var (
	topic     = "TrxState"
	redisAddr = "127.0.0.1:6379"
	ctx       = context.Background()
	logger    = watermill.NewStdLogger(false, false)
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
	repo := NewRepository(db)
	routers, err := NewRouters(ctx, &config, repo)
	if err != nil {
		panic(err)
	}
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	mux := NewHandler(repo, routers.CommandBus, routers.SSERouter)
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
