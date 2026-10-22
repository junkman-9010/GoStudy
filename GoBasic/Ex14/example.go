package main

import (
	"fmt"
)

// Case 1: Basic Variadic Function
// sumAll takes a variable number of integers and returns their sum.
func sumAll(nums ...int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum
}

// Case 2: Zero Arguments
// This is demonstrated in main(), showing that calling sumAll() returns 0.

// Case 3: Slice Unpacking
// demonstrateSliceUnpacking shows how to pass an existing slice to a variadic function using the '...' suffix.
func demonstrateSliceUnpacking() {
	nums := []int{10, 20, 30}
	fmt.Printf("Unpacking slice %v into sumAll: %d\n", nums, sumAll(nums...))
}

// Case 4: Mixed Parameters
// printWithPrefix shows that variadic parameters must be the final parameter in a function signature.
func printWithPrefix(prefix string, nums ...int) string {
	result := fmt.Sprintf("%s: ", prefix)
	for _, n := range nums {
		result += fmt.Sprintf("%d ", n)
	}
	return result
}

func main() {
	fmt.Println("--- Case 1: Basic Variadic Function ---")
	fmt.Printf("sumAll(1, 2, 3): %d\n", sumAll(1, 2, 3))

	fmt.Println("\n--- Case 2: Zero Arguments ---")
	fmt.Printf("sumAll(): %d\n", sumAll())

	fmt.Println("\n--- Case 3: Slice Unpacking ---")
	demonstrateSliceUnpacking()

	fmt.Println("\n--- Case 4: Mixed Parameters ---")
	fmt.Println(printWithPrefix("Numbers", 1, 2, 3, 4, 5))
}
