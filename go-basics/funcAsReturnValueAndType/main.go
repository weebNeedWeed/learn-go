package main

import (
	"fmt"
)

type testInt func(int) bool

func isOdd(a int) bool {
	return a%2 != 0
}

func isEven(a int) bool {
	return a%2 == 0
}

func filter(numbers []int, f testInt) (returnSlice []int) {
	for _, number := range numbers {
		if f(number) {
			returnSlice = append(returnSlice, number)
		}
	}
	return
}

func main() {
	originalSlices := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	oddSlices := filter(originalSlices, isOdd)
	evenSlices := filter(originalSlices, isEven)

	fmt.Println(oddSlices)
	fmt.Println(evenSlices)
}
