// https://golang.org/ref/spec#Handling_panics
// https://blog.golang.org/defer-panic-and-recover
package main

import (
	"fmt"
	// "runtime/debug"
)

// in the panic handling function, recover() must be directly called
// recover() returns the argument that was passed to the function that panicked
// if the return result is not nil, do the panic handling
// otherwise the follwing might be true
//     panic's argument was nil
//     the goroutine is not panicking
//     recover was not called directly by a deferred function
func D() {
	r := recover()
	if r != nil {
		fmt.Printf("recover(): %v\n", r)
		// debug.PrintStack()
	}
}

func panicFunc() {
	panic("the panic function")
}

func G() {
	// this is called after D()
	// i++
	defer func() {
		fmt.Println("deferred function called after D()")
		i++
	}()

	// handle any possible subsequent panic
	// must be handled in a deferred function call
	// the function could be literal (func(){}) or a named function, like D here
	defer D()

	// this is called before D() (after panic() but before recover())
	// however, it's not discarded but called normally
	defer func() {
		fmt.Println("deferred function called before D()")
		i++
	}()


	fmt.Println("G() function")

	// something bad happens
	// this could be a direct panic function call
	panicFunc()

	// statements after panic() and before recover() are discarded
	// except those in deferred function
	// all subsequent statements are discarded when the above panicFunc() panics
	fmt.Println("")
	i++
}

var i int

func main() {
	G()
	fmt.Println(i)
}
