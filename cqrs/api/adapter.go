package main

import (
	"encoding/json"
	"net/http"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
)

type orderplacedStreamAdapter struct {
	logger watermill.LoggerAdapter
}

func (f orderplacedStreamAdapter) InitialStreamResponse(w http.ResponseWriter, r *http.Request) (response interface{}, ok bool) {
	traceID := r.URL.Query().Get("traceid")
	if len(traceID) == 0 {
		return nil, false
	}
	return nil, true
}

func (f orderplacedStreamAdapter) NextStreamResponse(r *http.Request, msg *message.Message) (response interface{}, ok bool) {
	traceID := r.URL.Query().Get("traceid")
	var payload OrderPlaced
	err := json.Unmarshal(msg.Payload, &payload)
	if err != nil {
		return nil, false
	}
	if traceID == payload.RequestID {
		return payload, true
	}

	return nil, false
}
