package main

import (
	"github.com/ThreeDotsLabs/watermill/message"
)

func notifyMiddleware(pub message.Publisher) func(message.HandlerFunc) message.HandlerFunc {
	return func(next message.HandlerFunc) message.HandlerFunc {
		return func(msg *message.Message) ([]*message.Message, error) {
			msgs, err := next(msg)
			if err != nil {
				return msgs, err
			}
			err = pub.Publish(topic, msgs...)
			if err != nil {
				return nil, err
			}
			return msgs, nil
		}
	}
}
