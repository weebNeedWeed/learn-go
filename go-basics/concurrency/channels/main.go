package main

import "fmt"

func sum(nums []int, c chan int) {
	total := 0
	for _, value := range nums {
		total += value
	}
	c <- total
}

func main() {
	nums := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(nums[:len(nums)/2], c)
	go sum(nums[len(nums)/2:], c)

	x, y := <-c, <-c

	fmt.Printf("x = %d, y = %d; x + y = %d", x, y, x+y)
}
