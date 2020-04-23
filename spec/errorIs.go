// https://golang.org/ref/spec#Errors
// Go 2 Error Inspection
// https://go.googlesource.com/proposal/+/master/design/29934-error-values.md
// https://blog.golang.org/go1.13-errors (17 October 2019)
// https://golang.org/pkg/errors/
package main

import (
	"errors"
	"fmt"
	"strings"
)

type FatalError struct{ s string }

func (e FatalError) Error() string { return e.s }

// this Is method will be invoked by errors.Is when
// camparing a FatalError value to a target error value
// the return value of Is will be returned by errors.Is
func (FatalError) Is(target error) bool {
	return strings.ContainsAny(target.Error(), "Fatal")
}

func main() {
	// create an error chain
	uerr := errors.New("error 0")
	err := fmt.Errorf("error 1 wraps [%w]", uerr)

	// inspect wrapped error(s)
	// find error by value
	// the target error must be the identical error value (the same instance)
	targetErr := uerr
	match := errors.Is(err, targetErr)
	fmt.Printf("errors.Is(%q, %q) = %v\n", err, uerr, match)

	// anotherErr is a distinct error value even if the text is identical
	// see errors.New() doc for details
	anotherErr := errors.New("error 0")
	match = errors.Is(err, anotherErr)
	fmt.Printf("errors.Is(%q, %q) = %v\n", err, anotherErr, match)

	// customize the behavior of errors.Is
	tErr := errors.New("a sample fatal error")
	fErr := FatalError{"any error with (F/f)atal in text will match this error"}
	err = fmt.Errorf("error 1 wraps [%w]", fErr)
	match = errors.Is(err, tErr)
	fmt.Printf("errors.Is(%q, %q) = %v\n", err, tErr, match)
}
