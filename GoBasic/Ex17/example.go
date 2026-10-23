package main

import (
	"fmt"
)

// Case 3: Pointer Manipulation
// Function that modifies the value of an integer via its memory address.
func modifyValue(n *int) {
	fmt.Printf("  [Case 3] Inside function: modifying value at %p from %d to 100\n", n, *n)
	*n = 100
}

func main() {
	fmt.Println("--- Go Pointers & Address Operators Tutorial ---")

	// Case 1: Address Extraction
	// Using the '&' operator to find the memory address of a variable.
	var x int = 10
	fmt.Printf("Case 1 (Address Extraction):\n  Value: %d, Address: %p\n", x, &x)

	// Case 2: Dereferencing
	// Using a pointer variable and the '*' operator to access and modify the original value.
	var ptr *int = &x
	fmt.Printf("\nCase 2 (Dereferencing):\n  Pointer ptr points to: %p\n", ptr)
	fmt.Printf("  Original value via ptr: %d\n", *ptr)
	*ptr = 20
	fmt.Printf("  Value after dereference modification (*ptr = 20): %d\n", x)

	// Case 3: Pointer Manipulation
	// Passing a pointer to a function to alter the variable in the caller's scope.
	fmt.Printf("\nCase 3 (Pointer Manipulation):\n  Value before modifyValue: %d\n", x)
	modifyValue(&x)
	fmt.Printf("  Value after modifyValue: %d\n", x)

	// Case 4: Pointer Type Assignment
	// Assigning the address of one variable to another pointer variable.
	var y int = 50
	var ptrToY *int = &y
	fmt.Printf("\nCase 4 (Pointer Type Assignment):\n  Variable y: %d, Address: %p\n", y, &y)
	fmt.Printf("  ptrToY points to: %p\n", ptrToY)
	
	// Demonstrating update via the second pointer
	*ptrToY = 99
	fmt.Printf("  Value of y after modification via ptrToY: %d\n", y)
}
