package main

type PlaceOrderCommand struct {
	TrxId  int64 `json:"trxid"`
	UserId int64 `json:"user_id"`
}

type OrderPlaced struct {
	TrxId   int64 `json:"trxid"`
	OrderId int64 `json:"order_id"`
	UserId  int64 `json:"user_id"`
}

type TrxState struct {
	TrxId    int64  `json:"trxid"`
	Type     int    `json:"type"`
	State    int    `json:"state"`
	Name     string `json:"name"`
	Progress int    `json:"progress"`
}
