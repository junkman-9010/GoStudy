package main

import (
	"fmt"
)

func main() {
	// 1. Standard for loop (C-style): Counter from 1 to 5
	fmt.Println("--- Standard for loop (C-style) ---")
	for i := 1; i <= 5; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println("\n")

	// 2. While-style for loop (Condition-only)
	fmt.Println("--- While-style for loop ---")
	n := 1
	target := 5
	for n <= target {
		fmt.Printf("%d ", n)
		n++
	}
	fmt.Println("\n")

	// 3. Infinite loop with break
	fmt.Println("--- Infinite loop with break ---")
	count := 0
	for {
		count++
		fmt.Printf("%d ", count)
		if count >= 5 {
			break
		}
	}
	fmt.Println("\n")

	// 4. for range on a slice
	fmt.Println("--- for range on a slice ---")
	languages := []string{"Go", "Python", "Rust", "C++"}
	for index, lang := range languages {
		fmt.Printf("Index: %d, Value: %s\n", index, lang)
	}
	fmt.Println("\n")

	// 5. for range on a map
	fmt.Println("--- for range on a map ---")
	versions := map[string]string{
		"Go":     "1.21",
		"Python": "3.12",
		"Rust":   "1.75",
	}
	for name, ver := range versions {
		fmt.Printf("%s version: %s\n", name, ver)
	}
	fmt.Println("\n")

	// 6. continue usage: Skipping an even number
	fmt.Println("--- continue usage (skipping even numbers) ---")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("%d ", i)
	}
	fmt.Println()
}
