package main

import (
	"log"
)

func Infof(format string, v ...interface{})  { log.Printf("INFO: "+format, v...) }
func Warnf(format string, v ...interface{})  { log.Printf("WARN: "+format, v...) }
func Errorf(format string, v ...interface{}) { log.Printf("ERROR: "+format, v...) }
func Fatalf(format string, v ...interface{}) { log.Fatalf("FATAL: "+format, v...) }
