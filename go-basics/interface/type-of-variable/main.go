package main

import (
	"fmt"
	"strconv"
)

type List []Element
type Element interface{}

type Person struct {
	name string
	age  int
}

func (p Person) String() string {
	return "(name: " + p.name + " - age:    " + strconv.Itoa(p.age) + " years)"
}

func main() {
	list := make(List, 3)
	list[0] = 1
	list[1] = "Hello world"
	list[2] = Person{"Dennis", 70}

	for index, elm := range list {
		if value, ok := elm.(int); ok {
			fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
		} else if value, ok := elm.(string); ok {
			fmt.Printf("list[%d] is a string and its value is %s\n", index, value)
		} else if value, ok := elm.(Person); ok {
			fmt.Printf("list[%d] is a Person and its value is %v\n", index, value)
		}
	}

	fmt.Printf("\nUsing Switch\n")

	for index, elm := range list {
		switch value := elm.(type) {
		case int:
			fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
		case string:
			fmt.Printf("list[%d] is a string and its value is %s\n", index, value)
		case Person:
			fmt.Printf("list[%d] is a Person and its value is %v\n", index, value)
		default:
			fmt.Printf("Unknown type")
		}
	}
}
