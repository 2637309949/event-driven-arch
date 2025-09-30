package main

import (
	"math/big"
	"time"
)

type EventRaw struct {
	ID              int64     `db:"id" json:"id"`
	TxHash          string    `db:"tx_hash" json:"tx_hash"`
	LogIndex        int       `db:"log_index" json:"log_index"`
	BlockNumber     int64     `db:"block_number" json:"block_number"`
	BlockTime       time.Time `db:"block_time" json:"block_time"`
	ContractAddress string    `db:"contract_address" json:"contract_address"`
	EventSignature  string    `db:"event_signature" json:"event_signature"`
	EventName       string    `db:"event_name" json:"event_name"`
	Topic0          string    `db:"topic0" json:"topic0"`
	Topic1          string    `db:"topic1" json:"topic1"`
	Topic2          string    `db:"topic2" json:"topic2"`
	Topic3          string    `db:"topic3" json:"topic3"`
	Data            []byte    `db:"data" json:"data"`
	CreatedAt       time.Time `db:"created_at" json:"created_at"`
}

type EventParsed struct {
	ID              int64     `db:"id" json:"id"`
	TxHash          string    `db:"tx_hash" json:"tx_hash"`
	LogIndex        int       `db:"log_index" json:"log_index"`
	BlockNumber     int64     `db:"block_number" json:"block_number"`
	BlockTime       time.Time `db:"block_time" json:"block_time"`
	ContractAddress string    `db:"contract_address" json:"contract_address"`
	EventName       string    `db:"event_name" json:"event_name"`
	FromAddress     string    `db:"from_address" json:"from_address"`
	ToAddress       string    `db:"to_address" json:"to_address"`
	TokenID         *big.Int  `db:"token_id" json:"token_id"`
	Value           *big.Int  `db:"value" json:"value"`
	Metadata        string    `db:"metadata" json:"metadata"`
	CreatedAt       time.Time `db:"created_at" json:"created_at"`
}

type EventStats struct {
	ID             int64     `db:"id" json:"id"`
	EventLabel     string    `db:"event_label" json:"event_label"`
	EventName      string    `db:"event_name" json:"event_name"`
	EventCount     int       `db:"event_count" json:"event_count"`
	CreatedAt      time.Time `db:"created_at" json:"created_at"`
}
