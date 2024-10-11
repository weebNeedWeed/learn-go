package main

import "fmt"

func myFunc(args ...int) {
	for index, arg := range args {
		fmt.Printf("%v => %v\n", index, arg)
	}
}

func main() {
	myFunc(5, 2, 3, 4)
}
