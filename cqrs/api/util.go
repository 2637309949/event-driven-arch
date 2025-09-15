package main

import (
	"github.com/ThreeDotsLabs/watermill"
	"github.com/gin-gonic/gin"
)

func mustNew(r interface{}, err error) interface{} {
	if err != nil {
		panic(err)
	}
	return r
}

func mustCall(err error) {
	if err != nil {
		panic(err)
	}
}

func mustRoutine(fn func() error) {
	go func() {
		err := fn()
		if err != nil {
			panic(err)
		}
	}()
}

func newFakePlaceOrderCommand(userId int64) PlaceOrderCommand {
	return PlaceOrderCommand{
		UserId: userId,
	}
}

func requestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := watermill.NewUUID()
		c.Set("request_id", requestID)
		c.Next()
	}
}
