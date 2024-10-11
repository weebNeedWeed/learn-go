package main

import (
	"fmt"
	"os"
)

var user = os.Getenv("USER")

func initData() {
	if user == "" {
		panic("$USER is not found")
	}
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	initData()
}
