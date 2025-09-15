package main

import "time"

type Order struct {
	OrderId   int64
	UserId    int64
	State     int
	CreatedAt time.Time
}
