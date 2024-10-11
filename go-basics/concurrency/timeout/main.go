package main

import (
	"fmt"
	"time"
)

func main() {
	c, quit := make(chan int), make(chan int)

	go func() {
		for {
			select {
			case v := <-c:
				fmt.Println(v)
			case <-time.After(5 * time.Second):
				fmt.Println("timeout")
				quit <- 0
				break
			}
		}
	}()
	<-quit
}
