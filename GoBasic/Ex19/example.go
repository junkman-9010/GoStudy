package main

import (
	"fmt"
)

func main() {
	var nilPtr *int
	fmt.Println("--- Case 1: The Danger (Nil Dereference Panic) ---")
	// We use a deferred function to recover from the panic so the program doesn't just crash immediately
	// and can continue to show the other cases.
	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Recovered from panic: %v\n", r)
			}
		}()

		fmt.Printf("Attempting to dereference nil pointer: %v\n", nilPtr)
		_ = *nilPtr // This WILL cause a panic
	}()

	fmt.Println("\n--- Case 2: The Guard (Defensive Programming) ---")
	var safePtr *int = nil
	if safePtr == nil {
		fmt.Println("Guard check: Pointer is nil. Skipping dereference to prevent panic.")
	} else {
		fmt.Printf("Safe dereference: %d\n", *safePtr)
	}

	fmt.Println("\n--- Case 3: Safe Initialization ---")
	// Using new(int) allocates memory and returns a pointer to it
	initializedPtr := new(int)
	*initializedPtr = 42
	fmt.Printf("Initialized pointer value: %d\n", *initializedPtr)

	// Using address-of operator
	val := 100
	addrPtr := &val
	fmt.Printf("Pointer from address-of operator: %d\n", *addrPtr)

	fmt.Println("\n--- Case 4: Logging/Status ---")
	checkPointerStatus(nilPtr)
	checkPointerStatus(initializedPtr)
	checkPointerStatus(addrPtr)
}

func checkPointerStatus(ptr *int) {
	if ptr == nil {
		fmt.Printf("Status: Pointer is [NIL] (Address: %v)\n", ptr)
	} else {
		fmt.Printf("Status: Pointer is [NON-NIL] (Address: %v, Value: %d)\n", ptr, *ptr)
	}
}
