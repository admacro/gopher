// https://en.wikipedia.org/wiki/Variable_shadowing
// In computer programming, variable shadowing occurs when a variable declared
// within a certain scope (decision block, method, or inner class) has the
// same name as a variable declared in an outer scope. At the level of identifiers
// (names, rather than variables), this is known as name masking. This outer
// variable is said to be shadowed by the inner variable, while the inner
// identifier is said to mask the outer identifier. This can lead to confusion,
// as it may be unclear which variable subsequent uses of the shadowed variable
// name refer to, which depends on the name resolution rules of the language.

// More on Go's variable scope
// --> Go allows redeclaration <--
// Redeclaration does not introduce a new variable; it just assigns a new
// value to the original.
// An identifier declared in a block may be redeclared in an inner block.
// While the identifier of the inner declaration is in scope, it denotes
// the entity declared by the inner declaration.

package main

import (
	"fmt"
)

func Bar() (n int, err error) {
	return 123, fmt.Errorf("Invalid")
}

func Foo() (n int, err error) { // declare function with result parameters n and err
	if true {
		n, err := Bar()

		// compiler: result parameter n not in scope at return
		// compiler: result parameter err not in scope at return
		return // if no return here, there would be no shadowing as n and err are not used
	}
	return
}

// Go compiler detects and disallows some cases of shadowing.
// go build variable_shadowing.go
// # command-line-arguments
// ./variable_shadowing.go:35:3: n is shadowed during return
// ./variable_shadowing.go:35:3: err is shadowed during return
func main() {
	s, t := Foo()
	fmt.Println(s, t)
}
