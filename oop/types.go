// types
// https://golang.org/ref/spec#Types
// https://golang.org/ref/spec#Type_declarations
// https://golang.org/ref/spec#Type_identity
// https://yourbasic.org/golang/type-alias/
package main

import "fmt"

// type definition
// each type has an underlying type
// the underlying types of Go's predeclared types are the types themselves
// the underlying type of a defined type T is the underlying type to which
// T refers in its type declaration. This means the underlying types of all
// types go back to Go's predeclared

// the underlying type of string is string, and struct struct
// the underlying type of Human is struct
type Human struct {
	Sex string
	Age int
}

func (h *Human) Love() { fmt.Println("Human loves.") }

// the underlying type of Man is the underlying type of Human which is struct
// type Human and type Man are distinct types
type Man Human
type ManPointer *string

func (m ManPointer) Think() { fmt.Println("Man thinks.") }

// type alias
// type MaleHuman and type Man are identical types
type MaleHuman = Man

func main() {
	h := Human{Sex: "male", Age: 25}
	h.Love()

	man := Man{Sex: "gentleman", Age: 25}
	man.Think()

	var mh MaleHuman = man
	fmt.Printf("#%v\n", mh)
	mh.Think()
}
