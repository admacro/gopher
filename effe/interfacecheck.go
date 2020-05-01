// https://golang.org/doc/effective_go.html#blank_implements
// see ../spec/methodSet.go for more on method set
package main

import (
	"fmt"
)

type I interface {
	M() T
	// N() // add a new method to I, it won't compile (line#32,29)
}

type T struct {
	f string
}

// T implements I by implementing M
// provided I is not changed
// or else T is subject to possible failure of satisfying I
// thus, we need some mechanism to provent this from happening
// there are two ways to make sure the type of a value implements
// an interface: the static way and the dynamic way (see below)
func (t T) M() T { return t }

// the static way (hard check, must be guaranteed)
// static interface check
// performed at compile time
// the type of the return parameter is I, which must be
// implemented by the type of the return value t
// the following compiles if T (type of t) implements I
// otherwise type T does not satisfy the interface
// thus, the satisfaction of interfaces or the correctness of
// interface implementations is checked and guaranteed by compiler
func funcI(val interface{}) I {
	// the dynamic way (soft check, better be guaranteed)
	// runtime interface check by type assertion
	// performed at runtime
	// other examples:
	//     the (Unm/M)arshaler interface in encoding/json package
	//     the Stringer interface in fmt package
	if i, ok := val.(I); ok {
		fmt.Println(i, ok)
		t := i.M()
		return t // statec check here, t must implement I
	}
	fmt.Printf("%T does not implement interface I", val)
	return nil
}

// when no such code as above (usage of static conversion from T to I)
// is available, for whatever reason, there would be no way for compiler
// to verify statically if T implements I correctly
// yet we still want to make sure I is correctly implemented by T
// one way to do so is to assign a value of type T to a global variable of I using
// the blank identifier will make compiler to check the implementation correctness:
// if it compiles, T implements I; otherwise T doesn't implement I.
//
// if this compiles, T implements I, and as a result *T implements I as well
// because given a type T, the method set of *T is always the superset of or equals to
// the method set of T; thus, if T implements I, *T also implements I
// thus, the following line gaurantees that I is implemnted by both T and *T
var _ I = T{}

// if this compiles, *T implements I, but T doesn't implement I
// because methods with *T as receiver are not in the method set of T, which means
// the methods of I implemented by *T are not in the method set of T, therefore T
// does not implement I
// thus, the following line only gaurantees that I is implemnted by *T
var _ I = (*T)(nil)

func main() {
	funcI(T{"go"})
	funcI(123)
}
