package main

import (
	"fmt"
	"runtime"
)

func main() {
	go say("world")
	say("hello")
}

func say(s string) {
	for i := 1; i <= 5; i++ {
		runtime.Gosched()
		fmt.Println(s)
	}
}

// hello
// world
// hello
// world
// hello
// world
// hello
// world
// hello
