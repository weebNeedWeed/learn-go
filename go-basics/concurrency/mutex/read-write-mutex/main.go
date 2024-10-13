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
	m := sync.RWMutex{}

	go func() {
		m.RLock()
		defer m.RUnlock()

		nIsEven := isEven(n)

		time.Sleep(5 * time.Millisecond)

		if nIsEven {
			fmt.Printf("%d is even\n", n)
		} else {
			fmt.Printf("%d is odd\n", n)
		}
	}()

	go func() {
		m.RLock()
		defer m.RUnlock()
		nIsPos := n >= 0

		time.Sleep(5 * time.Millisecond)

		if nIsPos {
			fmt.Printf("%d is positive\n", n)
		} else {
			fmt.Printf("%d is negative\n", n)
		}
	}()

	go func() {
		m.Lock()
		n++
		m.Unlock()
	}()

	time.Sleep(time.Second)
}
