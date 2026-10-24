package main

import (
	"fmt"
)

// Case 2: Function via Pointer
// increment takes a pointer to an int and increments the value at that address.
func increment(n *int) {
	*n++
}

// Case 3: Pointer to Pointer
// updatePointer takes a pointer to a pointer and changes the address it points to.
func updatePointer(p **int, newAddr *int) {
	*p = newAddr
}

func main() {
	fmt.Println("--- Go Pointer Manipulation Patterns ---")

	// Case 1: Value Trace
	// Demonstrating how to track the identity of a variable vs its value.
	fmt.Println("\n[Case 1: Value Trace]")
	x := 10
	ptrX := &x
	fmt.Printf("[Initial] Address: %p, Value: %d\n", ptrX, *ptrX)

	// Modify value via pointer
	*ptrX = 20
	fmt.Printf("[Modified] Address: %p, Value: %d (Address remains unchanged)\n", ptrX, *ptrX)

	// Case 2: Function via Pointer
	fmt.Println("\n[Case 2: Function via Pointer]")
	y := 5
	fmt.Printf("[Before] Value: %d\n", y)
	increment(&y)
	fmt.Printf("[After]  Value: %d (Updated via pointer)\n", y)

	// Case 3: Pointer to Pointer
	fmt.Println("\n[Case 3: Pointer to Pointer]")
	z := 100
	w := 200
	ptrZ := &z
	ptrPtrZ := &ptrZ

	fmt.Printf("[Initial] ptrPtrZ points to ptrZ, which points to z. Value: %d, Address of z: %p\n", *ptrPtrZ, &z)

	// Change what ptrZ points to using ptrPtrZ
	updatePointer(ptrPtrZ, &w)

	fmt.Printf("[Updated] ptrPtrZ now points to ptrZ, which points to w. Value: %d, Address of w: %p\n", *ptrPtrZ, &w)

	// Case 4: Memory Mapping Visualization
	// A simple way to visualize the relationship.
	fmt.Println("\n[Case 4: Memory Mapping Visualization]")
	fmt.Println("Format: [Address] -> [Value]")
	fmt.Printf("%p -> %d\n", &x, x)
	fmt.Printf("%p -> %d\n", &y, y)
	fmt.Printf("%p -> %d\n", &z, z)
	fmt.Printf("%p -> %d\n", &w, w)
}
