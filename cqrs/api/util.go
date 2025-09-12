package main

import (
	"math/rand"
	"net/http"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/brianvoe/gofakeit/v6"
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

func wrapHttpHandler(h http.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		h(c.Writer, c.Request)
	}
}

func newFakePlaceOrderCommand(requestID string) PlaceOrderCommand {
	var products []Product

	for i := 0; i < rand.Intn(5)+1; i++ {
		products = append(products, Product{
			ID:   watermill.NewShortUUID(),
			Name: gofakeit.ProductName(),
		})
	}

	return PlaceOrderCommand{
		RequestID: requestID,
		Customer: Customer{
			ID:    watermill.NewULID(),
			Name:  gofakeit.Name(),
			Email: gofakeit.Email(),
			Phone: gofakeit.Phone(),
		},
		Address: Address{
			Street:  gofakeit.Street(),
			City:    gofakeit.City(),
			Zip:     gofakeit.Zip(),
			Country: gofakeit.Country(),
		},
		Products: products,
	}
}

func requestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := watermill.NewUUID()
		c.Set("request_id", requestID)
		c.Next()
	}
}
