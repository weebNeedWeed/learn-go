package main

import "fmt"

func demoError(a int) error {
	if a < 0 {
		return fmt.Errorf("a cannot be less than 0")
	}

	return nil
}

func main() {
	if err := demoError(-1); err != nil {
		fmt.Printf("An error occurred: %v", err)
	}
}
