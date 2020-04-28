// https://golang.org/ref/spec#Allocation
//
// new
// https://golang.org/pkg/builtin/#new
// https://golang.org/doc/effective_go.html#allocation_new
// The new built-in function allocates memory. The first argument is a
// type, not a value, and the value returned is a pointer to a newly allocated
// zero value of that type.
//
// make
// https://golang.org/pkg/builtin/#make
// https://golang.org/doc/effective_go.html#allocation_make
// The make built-in function allocates and initializes an object of type
// slice, map, or chan (only). Like new, the first argument is a type, not
// a value. Unlike new, make's return type is the same as the type of its
// argument, not a pointer to it.
//
// https://golang.org/ref/spec#Variables
// https://golang.org/ref/spec#Variable_declarations
// https://golang.org/ref/spec#The_zero_value
// https://golang.org/ref/spec#Assignments
package main

import "fmt"

func main() {
	// to understand new and make, variable, zero value, and assignment must be understood first
	// a variable is initialized to the zero value of the given type if no expression is given
	// if an expression is given, a variable is initialized to the computation result of the expression

	// the zero value of type string is "" (the empty string)
	// the following two are equivalent
	var s string       // a variable of type string implicitly initialized with zero value ""
	var sv string = "" // a variable of type string explicitly initialized with value ""
	fmt.Printf("(%T)(%q) == (%T)(%q) is %v\n", s, s, sv, sv, s == sv)

	// the zero value of pointer types is nil
	// *string is a pointer type, its zero value is nil
	// the following two are equivalent
	var sp *string        // a variable of type *string implicitly initialized with zero value nil
	var spv *string = nil // a variable of type *string explicitly initialized with value nil
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
	// the type of sn is *string
	// sn points to a value of type string, the value is ""
	var sn = new(string)
	fmt.Printf("new: (%T)(%p) --> %q\n", sn, sn, *sn)

	// slice
	// the following two are equivalent
	var ss []string        // a variable of type []string implicitly initialized with zero value nil
	var ssv []string = nil // a variable of type []string explicitly initialized with value nil
	fmt.Printf("%#v == nil is %v (length %d, capacity %d)\n", ss, ss == nil, len(ss), cap(ss))
	fmt.Printf("%#v == nil is %v (length %d, capacity %d)\n", ssv, ssv == nil, len(ssv), cap(ssv))

	// the following two are equivalent
	var ssp *[]string        // a variable of type *[]string implicitly initialized with zero value nil
	var sspv *[]string = nil // a variable of type *[]string explicitly initialized with value nil
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
	// ssn points to the zero value nil of type []string
	// the value ssn points to is nil, but ssn itself is not
	var ssn = new([]string)
	// slice: %p address of 0th element in base 16 notation,
	//        with leading 0x (https://golang.org/pkg/fmt/)
	fmt.Printf("new: %#v == nil is %v\n", *ssn, *ssn == nil)
	fmt.Printf("new: (%T)(%p) --> %#v\n", ssn, ssn, *ssn)
	fmt.Printf("new: (%T)(%p) == nil is %v\n", ssn, ssn, ssn == nil)

	// append to a nil slice
	// a nil slice has no underlying array with which it associates;
	// since a nil slice has length 0 and capacity 0, the appending
	// is beyond slice length and capacity, which will always create
	// a bigger array to hold extra elements, and a new slice points
	// to the newly allocated array; inefficient when the append
	// operation is constant (very often; frequently; repeatedly)
	// see ../xahlee/slice.go for more on append
	fmt.Printf("(length %d, capacity %d) %#v\n", len(ss), cap(ss), ss)
	for i := 0; i < 5; i++ {
		// when length and capacity is 0, append needs to create a new
		// underlying array and associate it to the slice, the length
		// and capacity is the number of items to append;
		// after that, the capacity dobules when append beyond the capacity
		// new capacity = current capacity * 2
		ss = append(ss, fmt.Sprintf("go%d", i))
		fmt.Printf("(length %d, capacity %d) %#v\n", len(ss), cap(ss), ss)
	}

	// thus, for efficiency reason, it's better to initialize a slice
	// with a reasonable length and capacity, using make

	// make
	// make creates slices, maps, and channels only, and it returns an
	// initialized (not zeroed) value of type T (not *T).
	// slices, maps and channels represent, under the covers, references to data
	// structures that must be initialized before use.
	// make initializes the internal data structure and prepares the value of these
	// three types for use
	//
	// make slice
	// see ../xahlee/slice.go for more examples
	var ssm = make([]string, 5, 5)
	fmt.Printf("make: %#v (length %d, capacity %d)\n", ssm, len(ssm), cap(ssm))
}
