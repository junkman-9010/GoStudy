package main

import (
	"fmt"
)

func main() {
	// 1. float64 to int (precision loss)
	// Converting a float to an int truncates the fractional part.
	var f64 float64 = 3.99
	i := int(f64)
	fmt.Printf("1. Float64 to Int: %f -> %d (Precision lost)\n", f64, i)

	// 2. int to int8 (overflow/truncation)
	// Converting a large int to a smaller type causes overflow.
	var largeInt int = 130
	i8 := int8(largeInt)
	fmt.Printf("2. Int to Int8: %d -> %d (Overflow: 130 wraps around in int8)\n", largeInt, i8)

	// 3. int to float64 (upcasting)
	// Converting an int to float64 is safe and preserves the value.
	var valInt int = 42
	f := float64(valInt)
	fmt.Printf("3. Int to Float64: %d -> %f (Safe upcasting)\n", valInt, f)
}
