package main

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/go-chi/chi/v5"
)

type HomeStats struct {
	Label string `json:"label"`
	Name  string `json:"name"`
	Value int    `json:"value"`
}

type Handler struct {
	*chi.Mux
	routers *Routers
}

func (h *Handler) Run(ctx context.Context) {
	ctx, _ = signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	server := &http.Server{Addr: ":8080", Handler: h}
	go h.routers.Run(ctx) // 确保注册完事件处理函数
	go func() {
		logger.Info("Server started at", watermill.LogFields{"port": "8080"})
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("ListenAndServe error", err, watermill.LogFields{})
		}
	}()
	<-ctx.Done()
	logger.Info("Shutting down server...", watermill.LogFields{})
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		logger.Error("ListenAndServe error", err, watermill.LogFields{})
	}
	logger.Info("Server exiting", watermill.LogFields{})
}

func NewHandler(repo *Repository, routers *Routers) *Handler {
	sseStream := occurredStreamAdapter{logger: logger, repo: repo}
	sseHandler := routers.SSERouter.AddHandler(topic, sseStream)
	r := chi.NewRouter()
	r.Handle("/*", http.FileServer(http.Dir("./view")))
	r.Route("/api", func(r chi.Router) {
		r.Get("/sse/occurred", sseHandler)
		r.Get("/sse/test", func(w http.ResponseWriter, r *http.Request) {
			is := ItemSet{}
			is.Key = "hello"
			is.Value = "2025"
			routers.EventBus.Publish(ctx, &is)
		})
	})
	h := Handler{}
	h.Mux = r
	h.routers = routers
	return &h
}

type occurredStreamAdapter struct {
	repo   *Repository
	logger watermill.LoggerAdapter
}

func (o occurredStreamAdapter) InitialStreamResponse(w http.ResponseWriter, r *http.Request) (response interface{}, ok bool) {
	o.logger.Info("occurredStreamAdapter.InitialStreamResponse", nil)
	eventStats, err := o.repo.QueryEventStats(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("eventStatsByName failed"))
		return nil, false
	}
	stats := []HomeStats{}
	for _, v := range eventStats {
		stats = append(stats, HomeStats{
			Label: v.EventLabel,
			Name:  v.EventName,
			Value: v.EventCount,
		})
	}
	return map[string]interface{}{
		"type": "stats",
		"data": stats,
	}, true
}

func (o occurredStreamAdapter) NextStreamResponse(r *http.Request, msg *message.Message) (response interface{}, ok bool) {
	var occurred EventOccurred
	err := json.Unmarshal(msg.Payload, &occurred)
	if err != nil {
		return nil, false
	}
	return map[string]interface{}{
		"type": "occurred",
		"data": HomeStats{
			Name:  occurred.EventName,
			Value: 1,
		},
	}, true
}
