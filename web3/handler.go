package main

import (
	"encoding/json"
	"net/http"

	"github.com/ThreeDotsLabs/watermill"
	watermillhttp "github.com/ThreeDotsLabs/watermill-http/v2/pkg/http"
	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/go-chi/chi/v5"
)

type HomeStats struct {
	Label string `json:"label"`
	Name  string `json:"name"`
	Value int    `json:"value"`
}

func NewHandler(repo *Repository, eventBus *cqrs.EventBus, sseRouter watermillhttp.SSERouter) *chi.Mux {
	sseStream := occurredStreamAdapter{logger: logger, repo: repo}
	sseHandler := sseRouter.AddHandler(topic, sseStream)
	r := chi.NewRouter()
	r.Handle("/*", http.FileServer(http.Dir("./views")))
	r.Route("/api", func(r chi.Router) {
		r.Get("/sse/occurred", sseHandler)
		r.Get("/sse/test", func(w http.ResponseWriter, r *http.Request) {
			is := ItemSet{}
			is.Key = "hello"
			is.Value = "2025"
			eventBus.Publish(ctx, &is)
		})
	})
	return r
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
