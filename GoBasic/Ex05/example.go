package main

import "fmt"

func main() {
	var i int
	var s string
	var b bool
	var f float64

	fmt.Printf("int:    %v (type: %T)\n", i, i)
	fmt.Printf("string: %q (type: %T)\n", s, s)
	fmt.Printf("bool:   %v (type: %T)\n", b, b)
	fmt.Printf("float64: %v (type: %T)\n", f, f)
}
