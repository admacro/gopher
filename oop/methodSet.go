// https://golang.org/ref/spec#Method_sets
package main

import "fmt"

type T struct {
	x, y int
}

func (t T) ReceiveValue() {
	fmt.Printf("ReceiveValue %#v\n", t)
	t.x++
	t.y++
}

func (t *T) ReceivePointer() {
	fmt.Printf("ReceivePointer %#v\n", t)
	t.x++
	t.y++
}

func newT() T {
	return T{}
}

func newPointerT() *T {
	return &T{}
}

func main() {
	// method set of type T includes all methods declared with receiver T: ReceiveValue()
	t := T{}
	t.ReceiveValue()
	fmt.Printf("T: %#v\n", t)

	// But, how can t call ReceivePointer()?
	// No, it can't! That's the truth. But why?
	// Because, AUTOMATICALLY!
	// Go compiler interprets t as &t when calling ReceivePointer() as the method has a pointer receiver;
	// also, t (the value), is passed to the method, not &t (the pointer).

	// Thus, the following two expressions are equivalent:
	t.ReceivePointer()
	(&t).ReceivePointer()
	fmt.Printf("T: %#v\n", t)

	// --------------------------------
	// And, ...
	// Similar things happen vice versa
	// --------------------------------

	// Method set of type *T includes all methods declared with receiver *T or T (that is,
	// it also contains the method set of T): Receivepointer() and ReceiveValue()
	pt := &T{}
	pt.ReceivePointer()
	fmt.Printf("T: %#v\n", pt)

	// But, how?
	// Again, AUTOMATICALLY!
	// Go compiler interprets pt as *pt when calling ReceiveValue() as the method has a value receiver;
	// also, *pt (the value), is passed to the method, not pt (the pointer).

	// Thus, the following two expressions are equivalent:
	pt.ReceiveValue()
	(*pt).ReceiveValue()
	fmt.Printf("T: %#v\n", pt)

	// --------------------------------
	// When method receiver is not a variable, Go compiler will not do the auto-interpretation above.
	// TODO Why? Because the receiver cannot be referenced by & or dereferenced by *?
	// --------------------------------

	// newT() returns a value of type T
	// Method set of type T includes promoted methods with receiver T, but not *T
	newT().ReceiveValue()
	// newT().ReceivePointer() // error: ReceiverPointer is not in method set of T

	// newPointerT() returns a value of type *T
	// Method set of type *T includes promoted methods with receiver T or *T
	newPointerT().ReceiveValue()
	newPointerT().ReceivePointer()
}
