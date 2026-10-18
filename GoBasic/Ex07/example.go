package main

import "fmt"

const (
	// Simple constants
	Pi      = 3.14159
	Version = "1.0.0"
)

const (
	// iota for basic enumeration (starts at 0)
	Pending = iota
	Active
	Deleted
)

const (
	// iota with arithmetic offset (starts at 1)
	StatusSuccess = iota + 1
	StatusFailure
	StatusError
)

func main() {
	fmt.Println("--- Simple Constants ---")
	fmt.Printf("Pi: %.5f\n", Pi)
	fmt.Printf("Version: %s\n", Version)

	fmt.Println("\n--- Iota Enumeration (Base 0) ---")
	fmt.Printf("Status Pending: %d\n", Pending)
	fmt.Printf("Status Active: %d\n", Active)
	fmt.Printf("Status Deleted: %d\n", Deleted)

	fmt.Println("\n--- Iota with Offset (Base 1) ---")
	fmt.Printf("Status Success: %d\n", StatusSuccess)
	fmt.Printf("Status Failure: %d\n", StatusFailure)
	fmt.Printf("Status Error: %d\n", StatusError)
}
