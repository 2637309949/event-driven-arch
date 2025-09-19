package main

import (
	"encoding/json"
	"net/http"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
)

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
