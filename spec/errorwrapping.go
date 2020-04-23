// https://golang.org/ref/spec#Errors
// https://blog.golang.org/error-handling-and-go (12 July 2011)
// https://golang.org/pkg/errors
package main

import (
	"errors"
	"fmt"
)

type ErrWrapper struct {
	s   string
	err error
}

// implement error interface
func (e ErrWrapper) Error() string { return fmt.Sprintf("wrapped error") }

// Unwrap method for errors.Unwrap(err)
// without this method, errors.Unwrap(err) returns nil for wrapped error values
func (e ErrWrapper) Unwrap() error { return e.err }

func main() {
	// create wrapped errors
	// to wrap an error, apply %w verb to the underlying error
	uerr := errors.New("error 0")
	err := fmt.Errorf("error 1 wraps [%w]", uerr)
	if err != nil {
		fmt.Println(err)
	}

	// unwrap an error to get the underlying error
	uerr1 := errors.Unwrap(err)
	if uerr1 != nil {
		fmt.Println(uerr1)
	}

	// more wrapping
	err = fmt.Errorf("error 2 wraps [%w]", err)
	if err != nil {
		fmt.Println(err)
	}
	err1 := errors.Unwrap(err)
	if err1 != nil {
		fmt.Println(err1)
	}

	// create custom wrapped errors
	w := ErrWrapper{"error wrapper", uerr}
	fmt.Printf("%#v\n", w)
	uerr2 := errors.Unwrap(w)
	if uerr2 != nil {
		fmt.Println(uerr2)
	}
}
