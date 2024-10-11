package main

import "fmt"

func main() {
	l, b := rectangle(10, 20)
	fmt.Printf("%v %v", l, b)
}

func rectangle(length uint, breadth uint) (perimeter uint, area uint) {
	perimeter = (length + breadth) * 2
	area = length * breadth
	return
}
