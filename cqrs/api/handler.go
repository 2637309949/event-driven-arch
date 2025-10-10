package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ThreeDotsLabs/watermill"
	watermillhttp "github.com/ThreeDotsLabs/watermill-http/v2/pkg/http"
	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/go-chi/chi"
)

func NewHandler(repo *Repository, commandBus *cqrs.CommandBus, sseRouter watermillhttp.SSERouter) *chi.Mux {
	sseStream := trxStreamAdapter{logger: logger, repo: repo}
	sseHandler := sseRouter.AddHandler(topic, sseStream)
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
			fmt.Println(placeOrderCommand.TrxId)
			err = commandBus.Send(ctx, placeOrderCommand)
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
	return r
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
