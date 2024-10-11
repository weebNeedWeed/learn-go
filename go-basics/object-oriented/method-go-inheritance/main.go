package main

import "fmt"

type human struct {
	name  string
	age   uint8
	phone string
}

type student struct {
	human
	school string
}

type employee struct {
	human
	company string
}

func (h human) sayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

// Overriding method
func (e employee) sayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
		e.company, e.phone) //Yes you can split into 2 lines here.
}

func main() {
	sam := employee{human{"Sam", 45, "111-888-XXXX"}, "Golang Inc"}
	mark := student{human{"Mark", 25, "222-222-YYYY"}, "MIT"}

	sam.sayHi()
	mark.sayHi()
}
