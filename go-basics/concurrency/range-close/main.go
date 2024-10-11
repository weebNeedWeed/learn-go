package main

import "fmt"

func fibonacci(n int, c chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- y
		x, y = y, x+y
	}
	close(c) // producer closes the channel(Do not close in consumer)
}

func main() {
	c := make(chan int)
	go fibonacci(10, c)
	for v := range c {
		fmt.Println(v)
	}
}
