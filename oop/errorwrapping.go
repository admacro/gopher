// https://golang.org/ref/spec#Errors
// https://blog.golang.org/error-handling-and-go (12 July 2011)
package main

import (
	"errors"
	"fmt"
)

func main() {
	// create wrapped errors
	// to wrap an error, apply %w verb to the underlying error
	uerr := errors.New("underlying error")
	err := fmt.Errorf("this error wraps [%w]", uerr)
	if err != nil {
		fmt.Println(err)
	}

	// unwrap an error to get the underlying error
	uerr1 := errors.Unwrap(err)
	if uerr1 != nil {
		fmt.Println(uerr1)
	}
}
