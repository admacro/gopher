// https://go.dev/tour/basics/16
package main

import "fmt"

func main() {
	// Numeric constants are high-precision values.
	// An untyped constant takes the type needed by its context.
	const Base = 1000
	const Rate = 0.1
	fmt.Printf("%T\n", Base)     // int
	fmt.Printf("%T\n", Base*2)   // int
	fmt.Printf("%T\n", Base*0.1) // float64 (type depends on its context, 0.1 is float64)

	var i float32 = 0.1
	fmt.Printf("%T\n", Rate*2) // float64 (type depends on its context, 0.1 is float64)
	fmt.Printf("%T\n", Rate*i) // float32 (type depends on its context, i is float32)

	var n = 1000
	fmt.Printf("%T\n", n*2) // int
	// fmt.Printf("%T\n", n * 0.1)	// constant 0.1 truncated to integer
	fmt.Printf("%T\n", float64(n)*0.1) // float64 (type must match)
}
