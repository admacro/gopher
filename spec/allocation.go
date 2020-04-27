// https://golang.org/ref/spec#Allocation
// https://golang.org/doc/effective_go.html#allocation_new
//
// https://golang.org/pkg/builtin/#new
// The new built-in function allocates memory. The first argument is a
// type, not a value, and the value returned is a pointer to a newly allocated
// zero value of that type.
//
// https://golang.org/pkg/builtin/#make
// The make built-in function allocates and initializes an object of type
// slice, map, or chan (only). Like new, the first argument is a type, not
// a value. Unlike new, make's return type is the same as the type of its
// argument, not a pointer to it.
//
// https://golang.org/pkg/builtin/#make
// https://golang.org/ref/spec#Variables
// https://golang.org/ref/spec#Variable_declarations
// https://golang.org/ref/spec#The_zero_value
// https://golang.org/ref/spec#Assignments
package main

import "fmt"

func main() {
	// to understand new and make, variable, zero value,
	// and assignment must be understood first
	// a variable is initialized to its zero value if no expression is given
	// if an expression is given, a variable is initialized
	// to the computation result of the expression

	// the zero value of type string is "" (the empty string)
	// the following two are equivalent
	var s string       // a variable of type string with zero value ""
	var sv string = "" // a variable of type string initialized with value ""
	fmt.Printf("(%T)(%q) == (%T)(%q) is %v\n", s, s, sv, sv, s == sv)

	// the zero value of pointer types is nil
	// *string is a pointer type, its zero value is nil
	// the following two are equivalent
	var sp *string        // a variable of type *string with zero value nil
	var spv *string = nil // a variable of type *string initialized with value nil
	// pointer: %p base 16 notation, with leading 0x (https://golang.org/pkg/fmt/)
	fmt.Printf("(%T)(%p) == (%T)(%p) is %v\n", sp, sp, spv, spv, sp == spv)

	// new
	// var p = new(T) is equivalent to the following:
	//     var t T
	//     var p *T = &t
	//
	// new string
	// the type passed to new is string
	// the zero value of type string is the empty string ""
	// new allocates memory for the zero value of type string
	// new creates a pointer pointing to the newly allocated zero value ""
	// new returns the pointer, its type is *string
	// the returned pointer is assgined to variable sn
	// the type of ss is *string
	// ss points to a value of type string, the value is ""
	var sn = new(string)
	fmt.Printf("new: (%T)(%p) --> %q\n", sn, sn, *sn)

	// slice
	// the following two are equivalent
	var ss []string        // a variable of type []string with zero value nil
	var ssv []string = nil // a variable of type []string initialized with zero value nil
	fmt.Printf("%#v == nil is %v\n", ss, ss == nil)
	fmt.Printf("%#v == nil is %v\n", ssv, ssv == nil)

	// the following two are equivalent
	var ssp *[]string        // a variable of type *[]string with zero value nil
	var sspv *[]string = nil // a variable of type *[]string with zero value nil
	fmt.Printf("(%T)(%p) == nil is %v\n", ssp, ssp, ssp == nil)
	fmt.Printf("(%T)(%p) == nil is %v\n", sspv, sspv, sspv == nil)

	// new slice
	// the type passed to new is []string (slice)
	// the zero value of type []string (slice) is nil
	// new allocates memory for the zero value of type []string
	// new creates a pointer pointing to the newly allocated zero value nil
	// new returns the pointer, its type is *[]string
	// the returned pointer is assgined to variable ssn
	// the type of ssn is *[]string
	// ssn points to the zero value of type []string, the value is nil
	// the value ssn points to is nil, but ssn itself is not
	var ssn = new([]string)
	// slice: %p address of 0th element in base 16 notation,
	//        with leading 0x (https://golang.org/pkg/fmt/)
	fmt.Printf("new: %#v == nil is %v\n", *ssn, *ssn == nil)
	fmt.Printf("new: (%T)(%p) --> %#v\n", ssn, ssn, *ssn)
	fmt.Printf("new: (%T)(%p) == nil is %v\n", ssn, ssn, ssn == nil)

	// slice created by new has no capacity (0)
	// you can't add items to it
	// to add items to slice, you need to use make
	fmt.Printf("new: length %d, capacity %d\n", len(*ssn), cap(*ssn))

	// make slice
	var ssm = make([]string, 5, 5)
	fmt.Printf("make: %#v == nil is %v\n", ssm, ssm == nil)
}
