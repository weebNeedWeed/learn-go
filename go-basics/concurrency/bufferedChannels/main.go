package main

import "fmt"

func main() {
	c := make(chan int, 2)
	c <- 2
	c <- 3
	fmt.Println(<-c, <-c)
}
