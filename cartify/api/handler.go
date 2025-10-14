package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/go-chi/chi"
)

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
	sseStream := trxStreamAdapter{logger: logger, repo: repo}
	sseHandler := routers.SSERouter.AddHandler(topic, sseStream)
	r := chi.NewRouter()
	r.Handle("/*", http.FileServer(http.Dir("./views")))
	r.Route("/api", func(r chi.Router) {
		r.Get("/trx/{id}", sseHandler)
		r.Post("/order", func(w http.ResponseWriter, r *http.Request) {
			var placeOrderCommand PlaceOrderCommand
			err := Decode(r.Body, &placeOrderCommand)
			if err != nil {
				http.Error(w, "invalid request body", http.StatusBadRequest)
				return
			}
			placeOrderCommand = newFakePlaceOrderCommand(placeOrderCommand.UserId)
			placeOrderCommand.TrxId = NextID()
			err = routers.CommandBus.Send(ctx, placeOrderCommand)
			if err != nil {
				http.Error(w, "invalid request body", http.StatusBadRequest)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			Encode(w, map[string]string{
				"status": "ok",
				"trxid":  strconv.FormatInt(placeOrderCommand.TrxId, 10),
			})
		})
	})
	h := Handler{}
	h.Mux = r
	h.routers = routers
	return &h
}

type trxStreamAdapter struct {
	logger watermill.LoggerAdapter
	repo   *Repository
}

func (f trxStreamAdapter) InitialStreamResponse(w http.ResponseWriter, r *http.Request) (response interface{}, ok bool) {
	id := chi.URLParam(r, "id")
	trxId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, false
	}
	trx, err := f.repo.TrxByID(r.Context(), trxId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, true
		}
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("TrxByID failed"))
		return nil, false
	}
	updated := TrxStateUpdated{}
	updated.Type = trx.Type
	updated.State = trx.Name
	updated.Progress = trx.Progress
	return updated, true
}

func (f trxStreamAdapter) NextStreamResponse(r *http.Request, msg *message.Message) (response interface{}, ok bool) {
	id := chi.URLParam(r, "id")
	trxId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, false
	}
	var trxState TrxState
	err = json.Unmarshal(msg.Payload, &trxState)
	if err != nil {
		return nil, false
	}
	if trxId == trxState.TrxId {
		updated := TrxStateUpdated{}
		updated.Type = trxState.Type
		updated.State = trxState.Name
		updated.Progress = trxState.Progress
		return updated, true
	}
	return nil, false
}
