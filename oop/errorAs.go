// https://golang.org/ref/spec#Errors
// Go 2 Error Inspection
// https://go.googlesource.com/proposal/+/master/design/29934-error-values.md
// https://blog.golang.org/go1.13-errors (17 October 2019)
// https://golang.org/pkg/errors/
package main

import (
	"errors"
	"fmt"
	"os"
	"reflect"
)

type OsPathLinkError struct{ s string }

func (e OsPathLinkError) Error() string { return e.s }

// this As method will be invoked by errors.As when
// camparing a FatalError value to a target error value
// the return value of As will be returned by errors.As
// and this As method is responsible for setting target
//
// OsError can represent all three errors in os package
// or in other words, all three errors can be treated AS OsError
func (e OsPathLinkError) As(target interface{}) (match bool) {
	val := reflect.ValueOf(target)
	switch target.(type) {
	case **os.PathError:
		pe := os.PathError{Op: "operation", Path: "path", Err: errors.New("os path error")}
		val.Elem().Set(reflect.ValueOf(&pe))
		match = true
	case **os.LinkError:
		le := os.LinkError{Op: "operation", New: "path", Old: "old", Err: errors.New("os link error")}
		val.Elem().Set(reflect.ValueOf(&le))
		match = true
	default:
		match = false
	}
	return
}

func main() {
	// create an error chain
	uerr := errors.New("error 0")
	err := fmt.Errorf("error 1 wraps [%w]", uerr)

	// find error by type
	// As finds an error that matches the target
	// for an error to match the target, the error value must be assginable
	// to what the target points to; therefore, the target must be a pointer
	// the type of the matching error must implement the error interface
	// when a match is found, As will assign the error value to the target
	//
	// target: ep (*error)
	// target points to: e (error)
	// when matching error is found, assgin the error value to e
	var e error
	ep := &e
	if errors.As(err, ep) {
		// if match, As will assign the match to ep
		fmt.Printf("found matching error: [%v]\n", *ep)
	}

	// custom error
	type NewError error
	ne := NewError(errors.New("new error"))
	err = fmt.Errorf("normal error wraps [%w]", ne)
	var nerr NewError
	nerrp := &nerr

	// target: nerrp (*NewError)
	// target points to: ce (NewError)
	// when matching error is found, assgin the error value to ce
	//
	// errors.As(err, nerrp) is equivalent to the following code (from errors.As source code):
	//     val := reflect.ValueOf(nerrp)
	//     val.Elem().Set(reflect.ValueOf(errors.Unwrap(err)))
	//     fmt.Printf("%T %v\n", val, val)
	// see errors.As source code for more details
	if errors.As(err, nerrp) {
		fmt.Printf("found matching error: [%v]\n", nerr)
	}

	// customize the behavior of errors.As
	oe := OsPathLinkError{"either os path or link error will match this error"}

	// target: pathErrPtr (**os.PathError)
	// target points to: pathErr (*os.PathError) (*os.PathError implements error interface)
	// when matching error is found, assgin the error value to pathErr
	//
	// match
	var pathErr *os.PathError
	pathErrPtr := &pathErr
	err = fmt.Errorf("normal error wraps [%w]", oe)
	if errors.As(err, pathErrPtr) {
		fmt.Printf("found matching error: [%v]\n", pathErr)
	}

	// same as above
	// match
	var linkErr *os.LinkError
	linkErrPtr := &linkErr
	err = fmt.Errorf("normal error wraps [%w]", oe)
	if errors.As(err, linkErrPtr) {
		fmt.Printf("found matching error: [%v]\n", linkErr)
	}

	// same as above
	// not match
	var sysCallErr *os.SyscallError
	sysCallErrPtr := &sysCallErr
	err = fmt.Errorf("normal error wraps [%w]", oe)
	if errors.As(err, sysCallErrPtr) {
		fmt.Printf("found matching error: [%v]\n", sysCallErr)
	} else {
		fmt.Printf("No matching error found for type: %T\n", sysCallErrPtr)
	}
}
