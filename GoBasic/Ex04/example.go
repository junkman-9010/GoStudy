package main

import "fmt"

// Package-level variable using 'var'
var packageVar string = "I am a package-level variable"

func main() {
	// Function-level variable using 'var' with explicit type
	var functionVarExplicit int = 42

	// Function-level variable using 'var' with type inference
	var functionVarInferred = 3.14

	// Function-level variable using short declaration (':=')
	shortDeclVar := "Hello, Go!"

	// Variable demonstrating the 'Zero Value' concept (declared but not initialized)
	var zeroValueVar int

	fmt.Println("Package-level variable:", packageVar)
	fmt.Println("Function-level (explicit type):", functionVarExplicit)
	fmt.Println("Function-level (type inference):", functionVarInferred)
	fmt.Println("Short declaration:", shortDeclVar)
	fmt.Println("Zero Value (int):", zeroValueVar)
}
