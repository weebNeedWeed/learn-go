package main

import (
	"fmt"
	"strconv"
)

type Human struct {
	name  string
	age   uint8
	phone string
}

func (h Human) String() string {
	return "Name:" + h.name + ",Age:" + strconv.Itoa(int(h.age)) + ",Phone:" + h.phone
}

func main() {
	h := Human{"Mark", 18, "333-222"}
	fmt.Print(h)
}
