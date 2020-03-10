package main

import (
	"fmt"
	"math"
)

// Go does not have classes.
// However, you can define methods on types.
type Vertex struct {
	x, y float64
}

// a method is a function with a special `receiver` argument
// the receiver appears between the func keyword and the method name
// in the case, the method Abs() has a receiver of type Vertex named v
// the return type of the method appears after the method name
//
// method must have exactly one `receiver`
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

// methods are functions
// a method is just a funcion with a receiver argement
//
// this is simpler in regular function syntax
// the functionality is the same as above
// but you need to call it in the form: Abs(v)
func Abs(v Vertex) float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

// methods can be defined with pointer receiver
// *T is the type of the receiver
// methods with pointer receivers can modify the value to wihch the receiver
// points
// it's very common to modify the receiver of a method, so pointer receivers
// are more common than value receivers which are just copies of the receivers
func (v *Vertex) Scale(f float64) {
	v.x = v.x * f
	v.y = v.y * f
}

// in regular function form
// not so convient when calling as you need to pass the receiver like this
// ScaleF(&v, f)
func ScaleF(v *Vertex, f float64) {
	v.x = v.x * f
	v.y = v.y * f
}

// value receiver
// Functions that take a value argument must take a value of that specific type
// when v.ScaleV(f) is called, Go copies v and pass it to the method
func (v Vertex) ScaleV(f float64) {
	v.x = v.x * f
	v.y = v.y * f
}

// methods can also be defined on non-struct types
// methods can only be declared with a receiver whose type is defined in
// the same package as the method
// this also excludes the built-in types, such as int, map, etc.
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	} else {
		return float64(f)
	}
}

func main() {
	var v = Vertex{3, 4}
	fmt.Println(v.Abs()) // 5
	fmt.Println(Abs(v))  // 5

	// As a convenience, Go interprets the statement v.Scale(f) as (&v).Scale(f)
	// since the Scale method has a pointer receiver.
	v.Scale(0.618) // value of v is taken as the receiver
	fmt.Println(v) // {1.854 2.472}

	(&v).Scale(0.618) // pointer can also be taken as the receiver
	fmt.Println(v)    // {1.145772 1.527696}

	v = Vertex{3, 4}
	// ScaleF(v, 0.618) // cannot use v (type Vertex) as type *Vertex in argument to ScaleF
	ScaleF(&v, 0.618) // must take a pointer, compare to above: v.Scale(f) and (&v).Scale(f)
	fmt.Println(v)    // {3, 4}

	// methods take value as receiver don't change the original
	v = Vertex{3, 4}
	v.ScaleV(0.618)
	fmt.Println(v) // {3, 4}

	// As a convenience, Go interprets the statement p.Scale(f) as (*p).Scale(f)
	p := &v
	p.ScaleV(0.618)
	fmt.Println(v) // {3, 4}

	var f MyFloat = 123.4
	fmt.Println(f.Abs())
	f = -123.4
	fmt.Println(f.Abs())
}
