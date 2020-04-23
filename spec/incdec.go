// https://golang.org/ref/spec#IncDec_statements
// The "++" and "--" statements increment or decrement their operands by the
// untyped constant 1. As with an assignment, the operand must be addressable
// or a map index expression. (map index expression is not addressable)
package main

import "fmt"

func main() {
	// integer
	a := 1
	a++
	fmt.Printf("a++ : %v\n", a)

	// floating-point
	pi := 3.1415926
	pi++
	fmt.Printf("pi++ : %v\n", pi)

	// rune
	r := 'R'
	r++
	fmt.Printf("r++ : %v\n", r)

	// string
	// st := "Abc"
	// st++ // Error: invalid operation: st++ (non-numeric type string)

	// slice indexing
	s := []int{1}
	s[0]--
	fmt.Printf("s[0]-- : %v\n", s[0])

	// map indexing
	m := map[int]int{1: 111}
	m[1]--
	fmt.Printf("m[1]-- : %v\n", m[1])
}
