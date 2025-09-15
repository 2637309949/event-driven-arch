package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
)

type trxStreamAdapter struct {
	logger watermill.LoggerAdapter
	repo   *Repository
}

func (f trxStreamAdapter) InitialStreamResponse(w http.ResponseWriter, r *http.Request) (response interface{}, ok bool) {
	id := r.PathValue("id")
	trxId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, false
	}
	trx, err := f.repo.TrxByID(r.Context(), trxId)
	if err != nil {
		return nil, false
	}
	updated := TrxStateUpdated{}
	updated.Type = trx.Type
	updated.State = trx.Name
	updated.Progress = trx.Progress
	return updated, true
}

func (f trxStreamAdapter) NextStreamResponse(r *http.Request, msg *message.Message) (response interface{}, ok bool) {
	id := r.PathValue("id")
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
