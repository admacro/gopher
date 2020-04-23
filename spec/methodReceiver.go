// https://golang.org/ref/spec#Method_declarations
// A method is a function with a receiver. A method declaration binds an identifier,
// the method name, to a method, and associates the method with the receiver's
// base type.
package main

import "fmt"

type T1 struct {
	PopularName string
}

// method receiver base type
// define PointerT1 as T1 pointer type
type PointerT1 *T1

// The method receiver parameter section must declare a single non-variadic parameter.
// the type of method receiver must be a defined type T or a pointer to a defined type T.
// T is called the receiver base type which cannot be a pointer or interface type

// PointerT1 is a defined type, but it's a pointer
// error: invalid receiver type PointerT1 (PointerT1 is a pointer type)
// func (pt PointerT1) MethodPt() { fmt.Println("test") }

// binds method MethodT1 with receiver type T1, to the base type T1
func (r T1) MethodT1(p int) { fmt.Printf("method receiver type is %T\n", r) }

// binds method MethodT1 with receiver type *T1, to the base type T1
func (r *T1) MethodT1Pointer() { fmt.Printf("method receiver type is %T\n", r) }

// method receiver identifier may be omitted if not used
func (*T1) NoReceiverReferencing() { fmt.Println("This method does not reference receiver.") }

// For a base type, the non-blank names of methods bound to it must be unique.
// If the base type is a struct type, the non-blank method and field names
// must be distinct.
// error: ./methodReceiver.go:32:6: type T1 has both field and method named PopularName
// func (T1) PopularName() string { return "" }

// the type of a method
// The type of a method is the type of a function with the receiver as first argument.
// For intance, the method MethodT1 defined above has type: func(r T1, p int)
// However, a function declared this way is not a method, nor a method of T1

// fakeMethod is not in the method set of T1
var fakeMethod = func(r T1, p int) {
	fmt.Println("This is a function, not a method.")
}

func main() {
	t1 := T1{}
	t1.MethodT1(123)
	t1.MethodT1Pointer()
	t1.NoReceiverReferencing()

	fakeMethod(t1, 123)
}
