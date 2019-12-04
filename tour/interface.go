package main

import (
	"fmt"
	"math"
)

// An interface type is a set of method signatures.

// A value of interface can hold any value that implements those methods.
// All methods of an interface a must be implemented in v when assigning v to a.

// When assigning v to a, v must have the same type as the receivers in all
// methods of a; in other words, v is the receiver of the methods in a.
// We can also say the type of v, T implements method m when T is same as the
// type of the receiver of method m.

// So, in general, to make things simple, all methods on a given type should
// have either value or pointer receivers, but not a mixture of both.

// Interfaces are implemented implicitly.
// This means a type implements an interface by implementing its methods.
// There is no such explicit declaration as "A implements B".

type Vertexer interface {
	Abs() float64

	// Error on line#26
	// cannot use &v (type *Vertex) as type Vertexer in assignment:
	// *Vertex does not implement Vertexer (missing Scale method)
	// Scale(f float64)
}

func main() {
	var vter Vertexer
	v := Vertex{3, 4}
	vter = &v
	fmt.Println(vter.Abs())
}

type Vertex struct {
	x, y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}
