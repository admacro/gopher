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

func main() {
	// method set of type T includes only methods with a value receiver: ReceiveValue()
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

	// And, ...
	// Similar things happen vice versa

	// method set of type *T includes only methods with a pointer receiver: ReceivePointer()
	pt := &T{}
	pt.ReceivePointer()
	fmt.Printf("T: %#v\n", pt)

	// But, how can t call ReceiveValue()?
	// No, it can't! That's the truth. But why?
	// Because, AUTOMATICALLY!
	// Go compiler interprets pt as *pt when calling ReceiveValue() as the method has a value receiver;
	// also, *pt (the value), is passed to the method, not pt (the pointer).

	// Thus, the following two expressions are equivalent:
	pt.ReceiveValue()
	(*pt).ReceiveValue()
	fmt.Printf("T: %#v\n", pt)

	// TODO
	// Easy, right?
	// Not so much when you know interface and the concept of addressable value
}

type S struct {
	T
}

func NewS() S {
	return S{}
}
