package main

import "time"

type Config struct {
	RedisAddr       string
	LogLevel        string
	ShutdownTimeout time.Duration
}

func NewConfig() *Config {
	return &Config{
		RedisAddr:       "localhost:6379",
		LogLevel:        "info",
		ShutdownTimeout: 10 * time.Second,
	}
}
