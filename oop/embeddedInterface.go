// https://golang.org/ref/spec#Interface_types
// Overlapping interfaces is illegal before Go version 1.14
// more at https://golang.org/doc/go1.14#language
//
// before Go#v1.14, compilation error:
// ./embeddedInterface.go:26:2: duplicate method x
// ./embeddedInterface.go:30:2: duplicate method b
// ./embeddedInterface.go:61:4: ambiguous selector cI.x
package main

import "fmt"

type A interface {
	a()
	x()
}

type B interface {
	b()
	x()
}

// The method set of T is the union of the method sets of T’s
// explicitly declared methods and of T’s embedded interfaces.

// The method set of C = c(), Cc(), and method sets of A and B
type C interface {
	// embedded interface
	A
	B

	// Overlapping interfaces is illegal before Go version 1.14
	// https://golang.org/doc/go1.14#language
	b()

	// methods with the same names must have identical signature
	// b(int) // illegal, error: b redeclared

	// C's explicitly declared methods (exported and non-exported)
	c()  // non-exported
	Cc() // exported
}

type Impl string

func (i *Impl) a()  { fmt.Println("a()") }
func (i *Impl) b()  { fmt.Println("b()") }
func (i *Impl) x()  { fmt.Println("x()") }
func (i *Impl) c()  { fmt.Println("c()") }
func (i *Impl) Cc() { fmt.Println("Cc()") }

func main() {
	var (
		aI A
		bI B
		cI C
	)
	fmt.Printf("%v\n", aI) // <nil>
	si := Impl("Implementation")
	aI = &si
	aI.x()
	bI = &si
	bI.x()
	cI = &si
	cI.x()
	fmt.Printf("%v", aI)
}

// An interface type T may not embed itself
// or any interface type that embeds T, recursively.
type Bad interface {
	// error: illegal cycle in declaration of Bad
	// Bad // self embedding
	// Worse // recursive self embedding
}

type Worse interface {
	Bad
}
