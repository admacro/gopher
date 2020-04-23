// types
// https://golang.org/ref/spec#Types
// https://golang.org/ref/spec#Type_declarations
// https://golang.org/ref/spec#Type_identity
// https://yourbasic.org/golang/type-alias/
package main

import (
	"fmt"
	"strconv"
)

// A defined type is always different from any other type, including
// the type it is created from

// defined type (from interface type)
type Intr interface {
	foo()
}

// NewIntr does not inherit any methods bound to the given type Intr,
// but the method set of NewIntr remains unchanged
// NewIntr has the same method set as Intr
type NewIntr Intr

// defined type (from a composite type: struct)
type Dt struct {
	name string
}

func (dt *Dt) String() string { return fmt.Sprintf("%#v", dt) }

// *Dt implements both Intr and NewIntr
// as both Intr and NewIntr have method foo() in their method sets
func (dt *Dt) foo() { fmt.Printf("%#v\n", dt) }

// A defined type does not inherit any methods bound to the given type
// NewDt is a defined type
// NewDt and Dt are two distinct types
// NewDt has the same fields as Dt, but the method set of NewDt is empty when it's defined
type NewDt Dt

// associate a new method with NewDt
func (ndt *NewDt) NewString() string { return fmt.Sprintf("%#v", ndt) }

// elements of composite type (struct, slice, array, map)
// method set of elements (Dt (embedded field), NewDt (normal field)) of
// composite type CompDt remain unchanged, and the method set of CompDt
// contains the method set of its embedded field Dt
type CompDt struct {
	Dt
	ndt NewDt
}

// type alias
// Et and Dt and the same type
type Et = Dt

// Type definitions may be used to define different boolean, numeric,
// or string types and associate methods with them
type OutOfBound bool
type Age int
type Pi float64

func (a Age) BirthYear() string {
	return strconv.Itoa(2020 - int(a))
}

func main() {
	dt := Dt{"Defined type"}
	fmt.Println(dt.String())

	// ERROR: ndt.String undefined (type NewDt has no field or method String)
	// ndt := NewDt{"Defined type"}
	// fmt.Println(ndt.String())

	// &dt can be assigned to a variabl of either Inter or NewIntr type
	var it Intr = &dt
	it.foo()
	var newIt NewIntr = &dt
	newIt.foo()

	cdt := CompDt{Dt: dt, ndt: NewDt{"Defined type"}}
	fmt.Println(cdt.ndt.NewString()) // methods of normal element

	// promoted methods of embedded field
	cdt.foo()
	fmt.Println(cdt.String())

	// define Age as int and associate a method BirthYear with it
	a := Age(23)
	fmt.Println(a.BirthYear())

	// assignability
	// https://golang.org/ref/spec#Assignability
	// b and type of 3.14 (float64) have identical underlying types (float64),
	// at float64 is not a defined type.
	var b Pi = 3.14
	fmt.Println(b)
}
