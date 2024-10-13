package main

import (
	"fmt"
	"time"
)

func isEven(n int) bool {
	return n%2 == 0
}

func main() {
	n := 0

	go func() {
		nIsEven := isEven(n)

		time.Sleep(50 * time.Millisecond)

		if nIsEven {
			fmt.Printf("%d is even\n", n)
		} else {
			fmt.Printf("%d is odd\n", n)
		}
	}()

	go func() {
		n++
	}()

	time.Sleep(time.Second)
}
