package main

import (
	"fmt"
	"sync"
	"time"
)

func isEven(n int) bool {
	return n%2 == 0
}

func main() {
	n := 0
	m := sync.Mutex{}

	go func() {
		m.Lock()
		defer m.Unlock()

		nIsEven := isEven(n)

		time.Sleep(50 * time.Millisecond)

		if nIsEven {
			fmt.Printf("%d is even\n", n)
		} else {
			fmt.Printf("%d is odd\n", n)
		}
	}()

	go func() {
		m.Lock()
		n++
		m.Unlock()
	}()

	time.Sleep(time.Second)
}
