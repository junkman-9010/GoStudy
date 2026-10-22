package main

import "fmt"

func main() {
	// Case 1: IIFE (Immediately Invoked Function Expression)
	func() {
		fmt.Println("Case 1: IIFE - This message is printed from an immediately invoked function.")
	}()

	fmt.Println()

	// Case 2: Function as Variable
	greet := func(name string) {
		fmt.Printf("Case 2: Function as Variable - Hello, %s!\n", name)
	}
	greet("Gopher")

	fmt.Println()

	// Case 3: Basic Closure
	counter := func() func() int {
		count := 0
		return func() int {
			count++
			return count
		}
	}()

	fmt.Println("Case 3: Basic Closure - Counter starts:")
	fmt.Printf("  Iteration 1: %d\n", counter())
	fmt.Printf("  Iteration 2: %d\n", counter())
	fmt.Printf("  Iteration 3: %d\n", counter())

	fmt.Println()

	// Case 4: Function Factory
	makeMultiplier := func(factor int) func(int) int {
		return func(n int) int {
			return n * factor
		}
	}

	double := makeMultiplier(2)
	triple := makeMultiplier(3)

	fmt.Println("Case 4: Function Factory - Multipliers:")
	fmt.Printf("  Double 5: %d\n", double(5))
	fmt.Printf("  Triple 5: %d\n", triple(5))
}
