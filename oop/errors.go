// https://golang.org/ref/spec#Errors
// https://blog.golang.org/error-handling-and-go (12 July 2011)
package main

import (
	"errors"
	"fmt"
	"time"
)

type InvalidZipcodeError struct {
	Zipcode string
}

func (err InvalidZipcodeError) Error() string {
	return fmt.Sprintf("Invalid zipcode: %s", err.Zipcode)
}

func main() {
	// a very basic error that only has a string message
	err := errors.New("something wrong happend")
	if err != nil {
		fmt.Println(err)
	}

	// error message formatting
	err = fmt.Errorf("something wrong happend at %v", time.Now())
	if err != nil {
		fmt.Println(err)
	}

	// custom error with more useful information
	// Invalidzipcodeerror implements error interface
	// thus a value of Invalidzipcodeerror can be assigned to a variable of error
	var e error = InvalidZipcodeError{"12345a"}
	if e != nil {
		fmt.Println(e)
	}
}
