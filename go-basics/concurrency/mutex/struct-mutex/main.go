package main

import (
	"fmt"
	"sync"
	"time"
)

type intLock struct {
	val int
	sync.Mutex
}

func (n *intLock) isEven() bool {
	return n.val%2 == 0
}

func main() {
	i := &intLock{
		val: 0,
	}

	go func() {
		i.Lock()
		defer i.Unlock()

		nIsEven := i.isEven()

		time.Sleep(5 * time.Millisecond)

		if nIsEven {
			fmt.Printf("%d is even\n", i.val)
		} else {
			fmt.Printf("%d is odd\n", i.val)
		}
	}()

	go func() {
		i.Lock()
		defer i.Unlock()
		i.val++
	}()

	time.Sleep(time.Second)
}
