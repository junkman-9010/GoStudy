package main

import (
	"errors"
	"fmt"
)

// Case 1: Basic - A simple function with a single return value.
func add(a, b int) int {
	return a + b
}

// Case 2: Multiple Returns - A function returning both quotient and remainder.
func divide(a, b int) (int, int) {
	return a / b, a % b
}

// Case 3: Error Pattern - A function returning (int, error) to handle negative input.
func factorial(n int) (int, error) {
	if n < 0 {
		return 0, errors.New("factorial is not defined for negative numbers")
	}
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result, nil
}

// Case 4: Slice Parameter - A function accepting a slice and returning its sum.
func sumSlice(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func main() {
	fmt.Println("--- Case 1: Basic ---")
	fmt.Printf("10 + 5 = %d\n", add(10, 5))

	fmt.Println("\n--- Case 2: Multiple Returns ---")
	q, r := divide(10, 3)
	fmt.Printf("10 / 3 = Quotient: %d, Remainder: %d\n", q, r)

	fmt.Println("\n--- Case 3: Error Pattern ---")
	res, err := factorial(5)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Factorial of 5 is: %d\n", res)
	}

	resNeg, errNeg := factorial(-1)
	if errNeg != nil {
		fmt.Printf("Factorial of -1 Error: %v\n", errNeg)
	} else {
		fmt.Printf("Factorial of -1 is: %d\n", resNeg)
	}

	fmt.Println("\n--- Case 4: Slice Parameter ---")
	nums := []int{1, 2, 3, 4, 5}
	fmt.Printf("Sum of %v = %d\n", nums, sumSlice(nums))
}
