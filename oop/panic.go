// https://golang.org/ref/spec#Handling_panics
// https://blog.golang.org/defer-panic-and-recover
// see recover.go for panic handling
package main

import "fmt"

// When the function F calls panic, execution of F stops, any
// deferred functions in F are executed normally
// same with deferred functions in fCaller and main
func F() {
	defer func() { fmt.Println("F() defered function") }()
	fmt.Println("F() statements before panic")
	panic("fatal error")
	fmt.Println("F() statements after panic")
}

func fCaller() {
	defer func() { fmt.Println("fCaller() defered function") }()
	F()
}

// the panicking terminating sequence:
// F() statements before panic
// F() defered function
// fCaller() defered function
// main() defered function
// panic: fatal error
func main() {
	defer func() { fmt.Println("main() defered function") }()
	fCaller()
}
