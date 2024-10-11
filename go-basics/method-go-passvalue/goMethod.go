package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type rectangle struct {
	length, breadth float64
}

func (c circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func (r rectangle) area() float64 {
	return r.length * r.breadth
}

func main() {
	r := rectangle{20, 10}
	c := circle{20}

	fmt.Println("Rectangle's area =", r.area())
	fmt.Println("Circle's area =", c.area())
}
