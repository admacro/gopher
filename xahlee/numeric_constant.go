package main

import "fmt"

func main() {
	// An untyped constant takes the type needed by its context.
	const Base = 1000
	fmt.Printf("%T\n", Base)     // int
	fmt.Printf("%T\n", Base*2)   // int
	fmt.Printf("%T\n", Base*0.1) // float64 (type depends on its context)

	var n = 1000
	fmt.Printf("%T\n", n*2) // int
	// fmt.Printf("%T\n", n * 0.1)	// constant 0.1 truncated to integer
	fmt.Printf("%T\n", float64(n)*0.1) // float64 (type must match)
}
