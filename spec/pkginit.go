// https://golang.org/ref/spec#Package_initialization
// Variables may also be initialized using functions named init declared
// in the package block, with no arguments and no result parameters.
package main

import "fmt"

// one or more init functions are allowed per package
// the init identifier can be used only to declare init functions
// yet init cannot be referred to from anywhere in a program
// which means init is reserved only for package initialization
// and you can't call them from anywhere in a program
func init() { fmt.Println("init variables") }

// a normal package-level variable (uninitialized)
var initialVar int

// a package-level variable initialized with a dummy value
// only one line is required
var limit = 1000
// var limit int = prepareLimit()	// error: prepareLimit() used as value

// the purpose of init functions is to make it possible to initialize
// variables that need to be prepared through multiple steps (statements)
func init() {
	// more operations need to be done to get a good value for limit
	// calculate limit
	// prepare limit
	// etc.
	fmt.Println("init limit")
	limit = 5479
}

// good idea but not allowed :-(
func prepareLimit() {
	limit = 904
}

// error: func init must have no arguments and no return values
// func init(int) int { return 1 }

func main() {
	fmt.Println("main() function")

	fmt.Println(limit)
	// init() // undefined: init
}
