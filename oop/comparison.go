// https://golang.org/ref/spec#Comparison_operators
package main

import "fmt"

func main() {
	// In any comparison, the first operand must be assignable to the type
	// of the second operand, or vice versa.
	var i = 2
	fmt.Println(i == 2.0)
	// fmt.Println(i == 3.14) // error: constant 3.14 truncated to integer

	// bool
	a, b := true, true
	fmt.Println(a == b)
	fmt.Println(a == !b)
	fmt.Println(!a == !b)

	// complex
	c, d := 123+1i, 123+1i
	fmt.Println(c == d)
	fmt.Println(real(c) == real(d))
	fmt.Println(imag(c) == imag(d))

	// string values are comparable and ordered, lexically byte-wise
	e, f, g := "ab", "af", "ca"
	fmt.Println(e == f, e < f, g > f)

	// pointers
	// two pointer values are equal if they point to the same variable or
	// if both are nil
	ep := &e
	ee := "ab"
	fmt.Println(&e == ep, &e == &ee)

	// channels
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch01 := ch1
	fmt.Println(ch1 == ch2)  // false
	fmt.Println(ch1 == ch01) // true

	// interface
	var i1, i2, i3 interface{} = 1, 1, 3.14
	fmt.Println(i1 == i2, i2 == i3)

	// interface and type
	var ii1, ii2 interface{} = "hello", 123
	var it string = "hello"
	fmt.Println(ii1 == it)    // true
	fmt.Println(ii2 == 123.0) // false

	// struct
	type St struct {
		string
		i interface{}
	}
	st1 := St{"hello", ii2}
	st2 := St{it, 123.0}
	fmt.Println(st1 == st2) // false

	// array
	// Two array values are equal if their corresponding elements are equal.
	a1 := [2]interface{}{st1, 456}
	a2 := [2]interface{}{St{it, 123}, 456}
	fmt.Println(a1 == a2) // true

	// a3 := [2]int{123, 456}
	// error: invalid operation: a1 == a3 (mismatched types [2]interface {} and [2]int)
	// fmt.Println(a1 == a3) // false

	// Slice, map, and function values are not comparable.
	// but they can be compared to nil
}
