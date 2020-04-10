// https://golang.org/ref/spec#Assignments
// LHS = RHS
// LHS must be addressable, a map index expression, or (for = assignments
// only) the blank identifier. Operands may be parenthesized.
package main

import "fmt"

func main() {
	// x op= y
	// is equivalent to: x = x op (y)
	// op is a binary arithmetic operator: +-*/% &|^ &^ << >>
	a, i, n := []int{1}, 0, 1
	a[i] <<= 2
	i &^= 1 << n
	fmt.Println(a[i], i)

	// tuple assignment
	var x, y int
	f := func() (int, int) { return 1, 1 }

	// form 1: RHS is a single multi-valued expression such as
	// a function call, a channel or map operation, or a type assertion
	x, y = f()
	fmt.Println(x, y)

	m := map[int]int{1: 123, 2: 456}
	var v int
	var ok bool
	v, ok = m[1]
	fmt.Println(v, ok)

	// form 2: RHS has the same number of expressions as LHS
	x, y = 123, 456
	fmt.Println(x, y)

	// blank identifier
	// provides a way to ignore right-hand side values in an assignment
	_ = a[i]
	x, _ = f()
	v, _ = m[1]
	var vv int
	vv = m[2] // same as vv, _ = m[2]
	fmt.Println(x, v, vv)

	// more tuple assginment
	var c, d int = 111, 999
	c, d = d, c
	fmt.Println(c, d)

	ss := []int{1, 2, 3}
	j := 0
	// first, LHS are evaluated from left to right, the result is: j, ss[0]
	// and both LHS and RHS expressions are evaluated in the usual order (see evalorder.go)
	// second, assignments are carried out in left-to-right order:
	//    set j = 1, then ss[0] = 2
	j, ss[j] = 1, 22
	fmt.Println(j, ss[j], ss) // here ss[j] is ss[1], so the output is: 1 2 [22 2 3]

	j = 0
	ss[j], j = 11, 1          // set ss[0] = 11, j = 1
	fmt.Println(j, ss[j], ss) // here ss[j] is ss[1], so the output is: 1 2 [11 2 3]

	// panic: runtime error: index out of range [3] with length 3
	// so, it's not a good practice to do tuple assginment if not absolutely necessary
	// as when there is an error, you can't tell immediately which assginment caused it
	i, j = 1, 3
	// ss[i], ss[j] = 4, 5 // error

	i = 2
	// this is different with: for i, v := range ss { break }
	for i, ss[i] = range ss { // set i, ss[2] = 0, ss[0]
		break
		// if no break
		// on the second iteration: set i, ss[0] = 1, ss[1]
		// on the third iteration: set i, ss[1] = 2, ss[2]
		// ss will be [2 11 11] at the end of the loop
	}
	fmt.Println(i, ss[i], ss)
}
