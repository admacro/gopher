// https://golang.org/ref/spec#Defer_statements
// A "defer" statement invokes a function whose execution is deferred to
// the moment the surrounding function returns, either because the surrounding
// function executed a return statement, reached the end of its function body,
// or because the corresponding goroutine is panicking.
//
// https://golang.org/doc/effective_go.html#defer
// Go's defer statement schedules a function call (the deferred function)
// to be run immediately before the function executing the defer returns.
// It's an unusual but effective way to deal with situations such as resources
// that must be released regardless of which path a function takes to return.
// The canonical examples are unlocking a mutex or closing a file.
package main

import "fmt"

func cleanup(s string) {
	fmt.Printf("cleaning up [%v]...\n", s)
}

func deferStacked() {
	fmt.Println("main program starts...")

	fmt.Println("some initial work")
	defer cleanup("mess of initial work")

	for i := 1; i <= 5; i++ {
		fmt.Printf("Process step %v\n", i)
		defer cleanup(fmt.Sprintf("mess of step %v", i))
	}

	fmt.Println("main program exists...")
}

// deferred functions are executed after any result parameters are set
// by that return statement but before the function returns to its caller.

// execution order
// 1. func literal value and its parameter is evaluated and saved
// 2. set result to 111
// 3. execute func literal with its parameter (by which result is changed)
// 4. the surrounding function deferFuncliteral returns (with result set to 234)
func deferFuncLiteral() (result int) {
	defer func(i int) {
		// result is accessed after it was set to 111 by the return statement
		result += i
	}(123)
	// return     // this returns 123
	return 111 // this returns 234
}

func main() {
	deferStacked()
	fmt.Println(deferFuncLiteral())
}

// defer func_name(args)
// defer is a new control flow, not in other popular languages
// but part of its purpose is similar to finally in Java:
//     make sure A is done when B is finished

// when defer is called, the args are evaluated immediately, but the
// function call is not executed until the surrounding function returns

// a very common use case is to make sure opened files are closed after
// the surrounding function exists by deferring file close right after
// opening/reading a file.

// Stacking defers
// defered functions are pushed onto a stack. When a function returns,
// its deferred calls are executed in last-in-first-out order.

// output of above program:
// main program starts...
// some initial work
// Process step 1
// Process step 2
// Process step 3
// Process step 4
// Process step 5
// main program exists...
// cleaning up [mess of step 5]...
// cleaning up [mess of step 4]...
// cleaning up [mess of step 3]...
// cleaning up [mess of step 2]...
// cleaning up [mess of step 1]...
// cleaning up [mess of initial work]...
