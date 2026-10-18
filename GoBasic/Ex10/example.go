package main

import (
	"fmt"
)

func main() {
	fmt.Println("--- 1. Basic Switch with Values ---")
	char := 'b'
	switch char {
	case 'a':
		fmt.Println("Found 'a'")
	case 'b':
		fmt.Println("Found 'b'")
	default:
		fmt.Println("Not a, b, or c")
	}

	fmt.Println("\n--- 2. Switch with Multiple Values ---")
	vowel := 'e'
	switch vowel {
	case 'a', 'e', 'i', 'o', 'u':
		fmt.Println(vowel, "is a vowel")
	default:
		fmt.Println(vowel, "is not a vowel")
	}

	fmt.Println("\n--- 3. Switch without Expression (Boolean-based) ---")
	num := 15
	switch {
	case num < 10:
		fmt.Println("Number is small")
	case num >= 10 && num < 20:
		fmt.Println("Number is medium")
	default:
		fmt.Println("Number is large")
	}

	fmt.Println("\n--- 4. Using Fallthrough ---")
	x := 1
	switch x {
	case 1:
		fmt.Println("Case 1 triggered")
		fallthrough
	case 2:
		fmt.Println("Case 2 triggered (via fallthrough)")
	case 3:
		fmt.Println("Case 3 triggered")
	}

	fmt.Println("\n--- 5. Type Switch ---")
	testTypes := []interface{}{"Hello, Go!", 42, 3.14, true}

	for _, val := range testTypes {
		fmt.Printf("Value: %v (Type: %T) -> ", val, val)
		switch v := val.(type) {
		case string:
			fmt.Printf("It's a string: %s\n", v)
		case int:
			fmt.Printf("It's an int: %d\n", v)
		case float64:
			fmt.Printf("It's a float: %f\n", v)
		case bool:
			fmt.Printf("It's a bool: %t\n", v)
		default:
			fmt.Println("Unknown type")
		}
	}
}
