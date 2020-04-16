// https://golang.org/ref/spec#Errors
// Go 2 Error Inspection
// https://go.googlesource.com/proposal/+/master/design/29934-error-values.md
// https://blog.golang.org/go1.13-errors (17 October 2019)
// https://golang.org/pkg/errors/
package main

import (
	"errors"
	"fmt"
)

type Address struct {
	Line1   string
	City    string
	Zipcode string
}

type InvalidZipcodeError struct {
	Zipcode string
	Err     error
}

func (err *InvalidZipcodeError) Error() string {
	return fmt.Sprintf("Invalid zipcode %s", err.Zipcode)
}

func (err *InvalidZipcodeError) Unwrap() error {
	return fmt.Errorf("Caused by: %w", err.Err)
}

type InvalidAddressError struct {
	Address
	InvalidZipcodeError error
}

func (err *InvalidAddressError) Error() string {
	return fmt.Sprintf("Invalid address %#v\n", err.Address)
}

func (err *InvalidAddressError) Unwrap() error {
	return fmt.Errorf("Caused by: %w\n", error(err.InvalidZipcodeError))
}

func validateAddress(a *Address) error {
	// validating address ...
	// address is not valid due to invalid zipcode
	return &InvalidAddressError{
		Address: *a,
		InvalidZipcodeError: &InvalidZipcodeError{
			a.Zipcode,
			errors.New("Invalid zipcode format"),
		},
	}
}

func main() {
	a := Address{
		Line1:   "2000 shoreline CT",
		City:    "Brisbane",
		Zipcode: "9412p",
	}
	err := validateAddress(&a)
	var invalidZipcodeError error
	if errors.As(err, &invalidZipcodeError) {
		fmt.Printf("validation failed: %+v\n", err)
	}
}
