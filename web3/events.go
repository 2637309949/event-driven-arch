package main

type ItemSetCommand struct {
	Key   [32]byte
	Value [32]byte
}

type ItemSet struct {
	Key   string
	Value string
}
