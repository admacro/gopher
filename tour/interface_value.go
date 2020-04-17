// interfaces
// https://tour.golang.org/methods/9
package main

import (
	"fmt"
)

type I interface {
	M()
}

type T struct {
	S string
}

// Error when there is no nil check when i.M() is called
// panic: runtime error: invalid memory address or nil pointer dereference
// [signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x109b0e6]
func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
	} else {
		fmt.Println(t.S)
	}
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func main() {
	// Under the hood, interface can be thought of as a tuple of a value and a
	// concrete type: (V, T),  V is a value of type T

	// An interface value holds a value of a specific underlying concrete type.
	// Calling a method on an interface value executes the method of the same name
	// on its underlying type.

	var i I
	describe(i) // (<nil> <nil>) value is nil, type is nil

	// interface value i holds a value of &{hello} of type *main.T
	var t = T{"hello"}
	i = &t
	describe(i) // (&{hello} *main.T)
	// i.M() executes M() of *main.T
	i.M() // hello

	// compare to above
	// same type, different value
	t = T{"world"}
	i = &t
	describe(i) // (&{world} *main.T)
	i.M()       // world

	var nt *T
	i = nt
	describe(i) // (<nil> *main.T)
	i.M()       // <nil>

	var f = F(123.45) // convert to F (type conversion)
	i = f
	describe(i) // (123.45 main.F)
	i.M()       // 123.45

	// nil interface value
	var ni I
	describe(ni) // (<nil> <nil>)

	// No type inside the interface tuple to indicate which concrete method to
	// call, thus runtime error:
	// panic: runtime error: invalid memory address or nil pointer dereference
	// [signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x109b660]
	// ni.M()
}

func describe(i interface{}) {
	if i == nil {
		fmt.Printf("nil (V: %v, T: %T)\n", i, i)
	} else {
		fmt.Printf("(V: %v, T: %T)\n", i, i)
	}
}
