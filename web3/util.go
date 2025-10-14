package main

import (
	stdSQL "database/sql"
)

func Open(driverName string, dataSourceName string) *stdSQL.DB {
	db, err := stdSQL.Open(driverName, dataSourceName)
	if err != nil {
		panic(err)
	}
	return db
}
