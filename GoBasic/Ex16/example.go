package main

import "fmt"

// Case 4: Typed function
// Defining a custom function type allows for clearer signatures and reusable patterns.
type Predicate func(int) bool

func main() {
	fmt.Println("--- Case 1: Function Assignment ---")
	// Case 1: Function assignment
	// Assigning an anonymous function to a variable.
	add := func(a, b int) int {
		return a + b
	}
	fmt.Printf("1 + 2 = %d\n", add(1, 2))

	fmt.Println("\n--- Case 2: Function as Argument ---")
	// Case 2: Function as argument
	// Passing a function (a predicate) as an argument to another function.
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	
	isEven := func(n int) bool {
		return n%2 == 0
	}
	
	filtered := filter(numbers, isEven)
	fmt.Printf("Even numbers: %v\n", filtered)

	fmt.Println("\n--- Case 3: Function as Return Value ---")
	// Case 3: Function as return value
	// A function that returns another function (a factory pattern).
	logger := createLogger("INFO")
	logger("This is an info message.")
	errorLogger := createLogger("ERROR")
	errorLogger("This is an error message.")

	fmt.Println("\n--- Case 4: Typed Function ---")
	// Case 4: Typed function usage
	// Using the 'Predicate' type defined above.
	isGreaterThanFive := Predicate(func(n int) bool {
		return n > 5
	})
	filteredLarge := filter(numbers, isGreaterThanFive)
	fmt.Printf("Numbers > 5: %v\n", filteredLarge)
}

// filter is a higher-order function that takes a Predicate and returns a slice of ints.
func filter(nums []int, p Predicate) []int {
	var result []int
	for _, n := range nums {
		if p(n) {
			result = append(result, n)
		}
	}
	return result
}

// createLogger returns a function that logs messages with a specific prefix.
func createLogger(prefix string) func(string) {
	return func(message string) {
		fmt.Printf("[%s] %s\n", prefix, message)
	}
}
