package main

import (
	"fmt"
)

// BasicDefer demonstrates a simple defer statement.
func BasicDefer() {
	fmt.Println("Entering BasicDefer")
	defer fmt.Println("Deferred: This runs at the end of BasicDefer")
	fmt.Println("Inside BasicDefer")
}

// LIFODemonstration demonstrates that defer statements are executed in Last-In, First-Out order.
func LIFODemonstration() {
	fmt.Println("\nEntering LIFODemonable")
	defer fmt.Println("Defer 1 (First in, Last out)")
	defer fmt.Println("Defer 2")
	defer fmt.Println("Defer 3 (Last in, First out)")
	fmt.Println("Inside LIFODemonstration")
}

// ArgumentEvaluation demonstrates that arguments to a deferred function are evaluated when the defer statement is executed.
func ArgumentEvaluation() {
	fmt.Println("\nEntering ArgumentEvaluation")
	x := 10
	defer func(val int) {
		fmt.Printf("Deferred function received captured value: %d\n", val)
	}(x) // x is evaluated here

	x = 20
	fmt.Printf("Value of x changed to: %d\n", x)
	fmt.Println("End of ArgumentEvaluation")
}

// ResourceCleanup simulates the closing of a resource using defer.
func ResourceCleanup() {
	fmt.Println("\nEntering ResourceCleanup")
	fmt.Println("Opening Resource...")
	defer fmt.Println("Resource Closed (via defer)")

	fmt.Println("Using Resource...")
}

func main() {
	fmt.Println("--- Starting Defer Patterns Demo ---")
	BasicDefer()
	LIFODemonstration()
	ArgumentEvaluation()
	ResourceCleanup()
	fmt.Println("\n--- Demo Completed ---")
}
