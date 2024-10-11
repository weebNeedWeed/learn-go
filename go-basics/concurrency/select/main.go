package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- y:
			x, y = y, x+y
		case <-quit:
			fmt.Println("Quit")
			return
			// default: // run when 2 channels are unavailable
			// 	fmt.Println("Blocking")
		}
	}
}

func main() {
	c, quit := make(chan int), make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()

	fibonacci(c, quit)
}
