package main

import "fmt"

func cleanup() {
	fmt.Println("cleaning up...")
}

func main() {
	defer cleanup()

	fmt.Println("main program...")
}

// defer func_name(args)
// defer is a new control flow, not in other popular languages
// but part of its purpose is similar to finally in Java:
//     make sure A is done when B is finished
// when defer is called, the args are evaluated, but not the function
// which will be called when the surrounding function exists
// a very common use case is deferring file close after opening/reading
// a file

// in the example above, cleanup() will be called when main() exits
// output:
// main program...
// cleaning up...
