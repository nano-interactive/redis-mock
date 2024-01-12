package main

import "fmt"

type Action struct {
	Key   string
	Value any
	Type  string
}

type RedisMock struct {
	data []Action
}

func main() {
	fmt.Println("Start Here...")
}
