package main

import (
	"encoding/json"
	"io"
	"time"
)

var (
	sf    *Snowflake
	epoch = time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC).UnixNano() / int64(time.Millisecond) // 例：epoch 设为 2020-01-01 00:00:00 UTC 的毫秒数
)

func init() {
	s, err := NewSnowflake(2, epoch)
	if err != nil {
		panic(err)
	}
	sf = s
}
func NextID() int64 {
	return sf.NextID()
}

func newFakePlaceOrderCommand(userId int64) PlaceOrderCommand {
	return PlaceOrderCommand{
		UserId: userId,
	}
}

func Decode(r io.Reader, v interface{}) error {
	err := json.NewDecoder(r).Decode(v)
	if err != nil {
		return err
	}
	return nil
}

func Encode(w io.Writer, v interface{}) error {
	return json.NewEncoder(w).Encode(v)
}
