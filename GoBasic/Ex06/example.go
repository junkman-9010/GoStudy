package main

import "fmt"

func main() {
	// Declaration of variables with basic types
	var myInt int = 42
	var myFloat float64 = 3.14159
	var myString string = "Hello, Go!"
	var myBool bool = true

	// Using fmt.Printf with specific verbs to demonstrate formatting
	fmt.Printf("Integer: %d\n", myInt)
	fmt.Printf("Float (2 decimal places): %.2f\n", myFloat)
	fmt.Printf("String: %s\n", myString)
	fmt.Printf("Boolean: %t\n", myBool)
	fmt.Printf("Default format (all): %v\n", myInt)
}
