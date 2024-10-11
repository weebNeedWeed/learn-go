package main

import "fmt"

var a = struct {
	name string
	age  int
}{"Mark", 12}

func main() {
	fmt.Println(a)
}
