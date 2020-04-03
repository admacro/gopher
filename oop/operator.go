// https://golang.org/ref/spec#Operators
// https://golang.org/ref/spec#Arithmetic_operators
package main

import "fmt"

func main() {
	// operator precedence
	// unary operators have the highest precedence
	// the unary operator bitwise complement ^ has precedence over binary operator XOR ^
	// ^x is m ^ x  with m = "all bits set to 1" for unsigned x, and m = -1 for signed x

	// unary and binary operator
	// the following two statements are equivalent
	// ^ is unary operator here, thus has precedence over binary operator <<
	fmt.Println("Results:", ^1<<2)
	fmt.Println("Results:", (^1)<<2)

	// binary operators
	// shift operators: << and >>, have precedence over bitwise operator ^
	fmt.Println("Results:", 0^1<<2)
	fmt.Println("Results:", 0^(1<<2))

	// Binary operators of the same precedence associate from left to right.
	fmt.Println("Results:", 1<<2>>2)
}
