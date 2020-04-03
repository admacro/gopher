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

	// division /, modulo %, right shift >>, and bitwise AND &
	// if a dividend is a non-negative, and the divisor is a
	// constant power of 2, then
	//     division can be replaced by a right shift operation
	//     remainder can be computed by a bitwise AND operation
	fmt.Println("Results:", 11/4)
	fmt.Println("Results:", 11>>2) // 2=4/2 (shift bits count = dividing-divisor / 2)
	fmt.Println("Results:", 11%4)
	fmt.Println("Results:", 11&3) // 3=4-1 (AND operand = dividing-divisor - 1)
}
