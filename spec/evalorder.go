// https://golang.org/ref/spec#Order_of_evaluation
// https://github.com/golang/go/issues/23188
// https://stackoverflow.com/questions/34232126/order-of-evaluation-in-a-slice-literal
package main

import (
	"fmt"
)

// when evaluating the operands of an expression, assignment, or return
// statement, all function calls, method calls, and communication operations
// are evaluated in lexical left-to-right order
// evaluation order of anything else other than what have mentioned above
// is not specified (by the spec), which means the compiler should decide the order
func main() {
	eval()

	a := 1
	f := func() int { a++; return a }
	g := func() int { a++; return a }

	// the evaluation order of a, f(), g() is different across implementations (compilers)
	// but the evaluation of f() and g() is ordered: f() is before g()
	// for the official Go implementation, the order seems to be: f() -> g() -> a
	// x may be [1, 2, 3], [2, 2, 3], [3, 2, 3] with other Go implementations
	x := []int{a, f(), g()}

	// same for the following
	// y may be {1: "good"} or {1: "bad"}, eval order between the two map assignments is not specified
	y := map[int]string{a: "good", a: "bad"}
	// y may be {1: 2} or {2: 2}, eval order between map key and value is not specified
	z := map[int]int{a: f()}

	fmt.Println(x, y, z)
}

func eval() {
	var x, y [2]int

	h := func() int { fmt.Println("h()"); return 1 }
	i := func() int { fmt.Println("i()"); return 1 }
	j := func() int { fmt.Println("j()"); return 1 }
	f := func() int { fmt.Println("f()"); return 1 }
	k := func() bool { fmt.Println("k()"); return true }
	g := func(int, int, int) int { fmt.Println("g()"); return 1 }

	var ok bool
	c := make(chan int)
	go func(ch chan int) { ch <- 1 }(c)

	// the function calls and communication happen in the order:
	//     f(), h(), i(), j(), <-c, g(), and k()
	// However, the order of those events compared to the evaluation
	// and indexing of x and the evaluation of y is not specified.
	// This means it's up to the compiler to decide the order of evaluation
	// of x and y, and indexing of x.
	y[f()], ok = g(h(), i()+x[j()], <-c), k()

	// the output does not have <-c, because there is no place to put
	// the printing statement, which, if is put inside the goroutine
	// above, like this:
	//     go func(ch chan int) { fmt.Println("<c"); ch <- 1 }(c)
	// there is no gurantee it will be executed in the right order
	// i.e. after j() and before g()
	fmt.Println(x, y, ok)
}
