package main

import "time"

type Trx struct {
	TrxId     int64  `json:"trxid"`
	Type      int    `json:"type"`
	State     int    `json:"state"`
	Name      string `json:"name"`
	Progress  int    `json:"progress"`
	CreatedAt time.Time
}
