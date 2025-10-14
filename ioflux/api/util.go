package main

import (
	stdSQL "database/sql"
	"encoding/json"
	"io"
	"mime/multipart"
	"path/filepath"
	"strings"
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

func DetectFileType(header *multipart.FileHeader) string {
	fileType := header.Header.Get("Content-Type")
	if fileType == "" {
		ext := strings.ToLower(filepath.Ext(header.Filename))
		switch ext {
		case ".jpg", ".jpeg":
			fileType = "image/jpeg"
		case ".png":
			fileType = "image/png"
		case ".gif":
			fileType = "image/gif"
		case ".txt":
			fileType = "text/plain"
		case ".csv":
			fileType = "text/csv"
		default:
			fileType = "application/octet-stream"
		}
	}
	return fileType
}

func Open(driverName string, dataSourceName string) *stdSQL.DB {
	db, err := stdSQL.Open(driverName, dataSourceName)
	if err != nil {
		panic(err)
	}
	return db
}
