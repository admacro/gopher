// https://golang.org/ref/spec#Operators
//
// Shift expression: 1<<33, i<<33, 1<<s
// The right operand in a shift expression must have integer type or be
// an untyped constant representable by a value of type uint. If the left
// operand of a non-constant shift expression is an untyped constant, it is
// first implicitly converted to the type it would assume if the shift expression
// were replaced by its left operand alone.
package main

import (
	"fmt"
)

func main() {
	var ii = 1.0<<33
	fmt.Printf("%T: %v", ii, ii)

	var s uint = 33

	// determine the type of left operand of a non-constant shift expression
	// 
	// first 1<<s is a non-constant shift expression
	// second, 1<<s is replaced by the left oprand of the shift expression, which is 1
	// now the whoe statement becomes: var i = 1
	// in the new statement, the type of 1 is the default type of constant 1, which is int
	// now, go back to the shift expression, we can know the type of the left operand
	// the type of it is int
	var i = 1<<s                  // 1 has type int

	// apply the same analysis to the following statement, you will know the type of the
	// left operand of the shift expression is int32, because after the replacement, the
	// new statement is: var j int32 = 1
	// thus the type of 1 is int32
	var j int32 = 1<<s            // 1 has type int32; j == 0

	// the same applies to the following statement
	var k = uint64(1<<s)          // 1 has type uint64; k == 1<<33
	var m int = 1.0<<s            // 1.0 has type int; m == 0 if ints are 32bits in size
	var n = 1.0<<s == j           // 1.0 has type int32; n == true
	var o = 1<<s == 2<<s          // 1 and 2 have type int; o == true if ints are 32bits in size
	var p = 1<<s == 1<<33         // illegal if ints are 32bits in size: 1 has type int, but 1<<33 overflows int
	var u = 1.0<<s                // illegal: 1.0 has type float64, cannot shift
	var u1 = 1.0<<s != 0          // illegal: 1.0 has type float64, cannot shift
	var u2 = 1<<s != 1.0          // illegal: 1 has type float64, cannot shift
	var v float32 = 1<<s          // illegal: 1 has type float32, cannot shift
	var w int64 = 1.0<<33         // 1.0<<33 is a constant shift expression
	var x = a[1.0<<s]             // 1.0 has type int; x == a[0] if ints are 32bits in size
	var a = make([]byte, 1.0<<s)  // 1.0 has type int; len(a) == 0 if ints are 32bits in size
}
