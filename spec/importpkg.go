// https://golang.org/ref/spec#Import_declarations
package main

import (
	// an explicit period (.) appears instead of a name,
	// all the package's exported identifiers declared in
	// that package's package block will be declared in
	// the importing source file's file block and must be
	// accessed without a qualifier.
	. "fmt"

	// package name specified as r
	r "math/rand"

	// package name omitted
	// defaults to the identifier specified in the
	// package clause of the imported package "time"
	"time"

	// blank identifier as explicit package name
	// this imports a package solely for its side-effects (initialization)
	_ "io"
)

func main() {
	r.Seed(time.Now().Unix())
	Println(r.Intn(1000))
}
