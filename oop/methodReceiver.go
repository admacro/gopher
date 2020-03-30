// https://golang.org/ref/spec#Method_declarations
// That method receiver parameter section must declare a single non-variadic parameter.
package main

import "fmt"

type T1 struct {
	x, y int
}

// method receiver base type
// define Tpointer as T pointer type
type PointerT1 *T1

// the type of method receiver must be a defined type T or a pointer to a defined type T.
// T is called the receiver base type which cannot be a pointer or interface type

// PointerT1 is a defined type, but it's a pointer
// error: invalid receiver type PointerT1 (PointerT1 is a pointer type)
// func (pt PointerT1) MethodPt() { fmt.Println("test") }

func (r T1) MethodT1()         { fmt.Printf("method receiver type is %T\n", r) }
func (r *T1) MethodT1Pointer() { fmt.Printf("method receiver type is %T\n", r) }

// method receiver identifier may be omitted if not used
func (*T1) NoReceiverReferencing() { fmt.Println("This method does not reference receiver.") }

func main() {
	t1 := T1{}
	t1.MethodT1()
	t1.MethodT1Pointer()
	t1.NoReceiverReferencing()
}
