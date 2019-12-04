// http://xahlee.info/golang/golang_closure.html
package main

import "fmt"

// Golang's function may be a closure.
// Closure is a programming language feature. It is a function such
// that all variables it uses will still work outside where it is defined.

// A closure is a function value that references variables from outside its
// body. The function may access and assign to the referenced variables; in
// this sense the function is "bound" to the variables.
// Each closure is bound to its own variables.

// For example:
//   1. a function f defines a local variable a and a local function g
//   2. a is used in g
//   3. f returns g
//   4. f is called and result is assigned to h
//   5. h is now a function that is g. all variables used by h, including a
//  still work, even though they were local variables of f
//   6. h is called a closure

// In practice, closure means you can create a function that maintains a state,
// without using global variable.

// You create a closure by defining a function that returns a function. Then
// call the function and assign the result to a variable. The variable is a new
// function, a closure.

// Closure is basically the core concept of "class and object" in object-oriented
// programming languages, like Java, Python, Ruby, etc.

func main() {
	var class = func(x int) (getter func() int, setter func(x int)) {
		var a = x
		getter = func() int {
			return a
		}
		setter = func(x int) {
			a = x
		}
		return getter, setter
	}

	var get, set = class(1)
	fmt.Printf("get(): %v\n", get()) // get(): 1
	fmt.Println("set(3)")
	set(3)
	fmt.Printf("get(): %v\n", get()) // get(): 3

	var read, write = class(100)
	fmt.Printf("read(): %v\n", read()) // read(): 100
	fmt.Println("write(3000)")
	write(3000)
	fmt.Printf("read(): %v\n", read()) // read(): 3000

}
