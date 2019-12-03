package main

import "fmt"

func cleanup(s string) {
	fmt.Printf("cleaning up [%v]...\n", s)
}

func main() {
	fmt.Println("main program starts...")

	fmt.Println("some initial work")
	defer cleanup("mess of initial work")

	for i:= 1; i <= 5; i++ {
		fmt.Printf("Process step %v\n", i)
		defer cleanup(fmt.Sprintf("mess of step %v", i))
	}

	fmt.Println("main program exists...")
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

