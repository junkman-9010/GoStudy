package main

import "fmt"

func main() {
	// 1. Basic if-else
	number := -5
	if number >= 0 {
		fmt.Printf("%d is non-negative\n", number)
	} else {
		fmt.Printf("%d is negative\n", number)
	}

	// 2. If-else if-else chain
	score := 85
	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B")
	} else if score >= 70 {
		fmt.Println("Grade: C")
	} else {
		fmt.Println("Grade: F")
	}

	// 3. If statement with an initialization component
	// The variable 'err' is scoped only to the if/else block
	if err := checkStatus(); err != nil {
		fmt.Printf("Error encountered: %v\n", err)
	} else {
		fmt.Println("Status check successful!")
	}
}

// checkStatus is a dummy function to demonstrate the 'if with init' pattern
func checkStatus() error {
	// Simulating a successful check
	return nil
}
