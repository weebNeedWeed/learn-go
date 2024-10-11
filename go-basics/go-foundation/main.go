package main

import "fmt"

// define by group
var (
	hello string
	world string
)

const (
	x1 = 1
	x2 = 2
)

const (
	y1 = iota // y1 = 0
	y2        // y2 = 1
	y3        // y3 = 2
	y4        // y4 = 3
)

func main() {
	// copyDemo()

	// errorHandlingForMapDemo()

	stringToByteDemo()
}

func copyDemo() {
	defaultBytes := []byte{'a', 'b', 'c', 'd'}

	// must be the same as defaultBytes's length
	copiedBytes := make([]byte, len(defaultBytes))
	copy(copiedBytes, defaultBytes)

	fmt.Println(copiedBytes)
}

func errorHandlingForMapDemo() {
	exampleMap := make(map[string]string)
	exampleMap["hello"] = "world"

	if _, ok := exampleMap["hello1"]; ok == false {
		fmt.Printf("hello1 is missing")
	}
}

func stringToByteDemo() {
	exampleString := "hello world"
	exampleBytes := []byte(exampleString)
	fmt.Println(exampleBytes)
}
