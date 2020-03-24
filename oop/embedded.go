// https://golang.org/ref/spec#Struct_types
package main

import "fmt"

type T struct {
	x, y int
}

func (t *T) ReceivePointerT() {
	fmt.Printf("ReceivePointer: %#v\n", t)
}

func (t T) ReceiveValue() {
	fmt.Printf("ReceiveValue: %#v\n", t)
}

func (t T) IdenticalMethod() {
	fmt.Printf("IdenticalMethod of T: %#v\n", t)
}

type Pt struct {
	y, z int
}

func (pt Pt) ReceiveValuePt() {
	fmt.Printf("ReceiveValuePt: %#v\n", pt)
}

func (pt *Pt) ReceivePointer() {
	fmt.Printf("ReceivePointer: %#v\n", pt)
}

func (pt Pt) IdenticalMethod() {
	fmt.Printf("IdenticalMethod of Pt: %#v\n", pt)
}

type S struct {
	// embedded fields
	T   // embedded value type
	*Pt // embedded pointer type (Pt must not be an interface type or a pointer type)
}

func main() {
	s := &S{
		Pt: &Pt{}, T: T{},
	}
	fmt.Printf("%#v\n", s) // main.S{T:main.T{x:0, y:0}, Pt:(*main.Pt)(0xc0000140a0)}

	// A field or method f of an embedded field in a struct x is called promoted if
	// x.f is a legal selector that denotes that field or method f.

	// for f in x.f, if there is not exactly one f with shallowest depth, the selector expression is illegal.
	// see selector rules #1 at https://golang.org/ref/spec#Selectors

	// promoted fields
	fmt.Printf("s.x = %v\n", s.x)
	fmt.Printf("s.z = %v\n", s.z)

	// fmt.Println(s.y) // error: ambiguous selector s.y
	fmt.Printf("s.T.y = %v\n", s.T.y)
	fmt.Printf("s.Pt.y = %v\n", s.Pt.y)

	// promoted methods
	s.ReceiveValue()
	s.ReceiveValuePt()
	s.ReceivePointer()
	s.ReceivePointerT()

	// s.IdenticalMethod() // error: ambiguous selector s.IdenticalMethod
	s.T.IdenticalMethod()
	s.Pt.IdenticalMethod()
}
