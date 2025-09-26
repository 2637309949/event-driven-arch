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

func NewHandler(repo *Repository, eventBus *cqrs.EventBus, sseRouter watermillhttp.SSERouter) *chi.Mux {
	sseStream := itemStreamAdapter{logger: logger}
	sseHandler := sseRouter.AddHandler(topic, sseStream)
	r := chi.NewRouter()
	r.Route("/api", func(r chi.Router) {
		r.Get("/sse/event", sseHandler)
		r.Get("/sse/test", func(w http.ResponseWriter, r *http.Request) {
			is := ItemSet{}
			is.Key = "hello"
			is.Value = "2025"
			eventBus.Publish(ctx, &is)
		})
	})
	return r
}

type itemStreamAdapter struct {
	logger watermill.LoggerAdapter
}

func (f itemStreamAdapter) InitialStreamResponse(w http.ResponseWriter, r *http.Request) (response interface{}, ok bool) {
	return "", true
}

func (f itemStreamAdapter) NextStreamResponse(r *http.Request, msg *message.Message) (response interface{}, ok bool) {
	var itemSet ItemSet
	err := json.Unmarshal(msg.Payload, &itemSet)
	if err != nil {
		return nil, false
	}
	return itemSet, true
}
