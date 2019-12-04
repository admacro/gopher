package main

import (
	"fmt"
)

// Empty interface specifies zero methods.
// An empty interface may hold values of any type. (Every type implements at
// least zero methods.)

// Empty interfaces are used by code that need to handle values of unknown type.
// For example, fmt.Print takes any number of arguments of type interface{}.
func main() {
	var i interface{}
	describe(i) // (<nil> <nil>)

	i = 123
	describe(i) // (123 int)

	i = "hello"
	describe(i) // (hello string)
}

func describe(i interface{}) {
	fmt.Printf("(%v %T)\n", i, i)
}
