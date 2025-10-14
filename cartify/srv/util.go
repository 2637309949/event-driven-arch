package main

import (
	stdSQL "database/sql"
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

func Open(driverName string, dataSourceName string) *stdSQL.DB {
	db, err := stdSQL.Open(driverName, dataSourceName)
	if err != nil {
		panic(err)
	}
	return db
}
