// https://golang.org/ref/spec#Method_expressions
// If M is in the method set of type T, T.M is a function that is callable
// as a regular function with the same arguments as M prefixed by an additional
// argument that is the receiver of the method.
package main

import "fmt"

type T struct{ x int }

func (t T) Mv(a int) int {
	v := a + t.x
	fmt.Println(v)
	return v
}

func (t *T) Mp(b float64) float64 {
	v := b + float64(t.x)
	fmt.Println(v)
	return v
}

type Intr interface {
	Minter(s string) string
}

func (t T) Minter(s string) string {
	ss := s + "!"
	fmt.Println(ss)
	return ss
}

func main() {
	t := T{123}
	fmt.Println(t.Mv(111), t.Mp(3.14))

	// T.Mv is a method expression
	// T.Mv yields a function equivalent to Mv but with
	// an explicit receiver as its first argument
	// the yielded function has thes signature: func(t T, a int) int
	// the function may be called normally with an explicit receiver

	// the following five are equivalent
	t.Mv(222)      // 1
	T.Mv(t, 222)   // 2
	(T).Mv(t, 222) // 3

	fv := T.Mv // 4
	fv(t, 222)

	fp := (T).Mv // 5
	fp(t, 222)

	// same with method with a pointer receiver
	// (*T).Mp yields: func(t *T, b float64) float64
	// (*T).Mv yields: func(t *T, a int) int <----- Note the pointer receiver
	t.Mp(0.618)
	(&t).Mp(0.618)
	(*T).Mp(&t, 0.618)

	// although &t is passed, the underlying method gets
	// the value of t by indirection: *(&t)
	(*T).Mv(&t, 333)

	// interface
	// derive function from a method of an interface type
	var intr Intr = &t

	// the following are equivalent
	intr.Minter("hello")       // 1
	Intr.Minter(intr, "hello") // 2
	im := Intr.Minter          // 3
	im(intr, "hello")
}
