// https://golang.org/ref/spec#Function_literals
// A function literal represents an anonymous function
// Function literals are closures: they may refer to variables defined
// in a surrounding function. Those variables are then shared between the
// surrounding function and the function literal, and they survive as long
// as they are accessible.
// see ../xahlee/closure.go for more on closure
package main

import "fmt"

// this is a function type
type Jota func() int

func makeJota() Jota {
	var v int

	// assgin a function literal to f
	f := func() int {
		v = v + 1
		return v
	}
	return f
}

func main() {
	jota := makeJota()
	fmt.Println(jota())
	fmt.Println(jota())
	fmt.Println(jota())
	fmt.Println(jota())
}
