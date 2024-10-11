package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 4.65

	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)

	fmt.Println("type:", t)
	fmt.Println("value:", v)
	fmt.Println("kind is float64:", v.Kind() == reflect.Float64)

	// Change value of variable
	p := reflect.ValueOf(&x) // *float64
	v2 := p.Elem()
	v2.SetFloat(7.1)
	fmt.Println("value of x now:", x)
}
