// https://golang.org/ref/spec#Calls
package main

import (
	"fmt"
	"math"
)

func coordinate(a, b float64) (float64, float64) {
	return a, b
}

func C(a, b float64) float64 {
	return math.Sqrt(a*a + b*b)
}

func Point3D(x, y, z float64) (float64, float64, float64) {
	return x, y, z
}

func Cc(x, y float64, zz ...float64) float64 {
	zzLen := len(zz)
	if zzLen == 0 {
		return math.Sqrt(x*x + y*y)
	} else if zzLen == 1 {
		z := zz[0]
		return math.Sqrt(x*x + y*y + z*z)
	}
	return 0
}

func main() {
	a, b := 3.0, 4.0

	// the return values of coordinate(a, b) are assigned to
	// the parameters of function C in order
	c := C(coordinate(a, b))
	fmt.Printf("%v*%v + %v*%v = %v*%v\n", a, a, b, b, c, c)

	// the two return values of coordinate(a, b) are assgined to
	// the first two parameters of function Cc in order, leaving
	// the third zz blank as zz is an optionl parameter
	d := Cc(coordinate(a, b))
	fmt.Printf("%v*%v + %v*%v = %v*%v\n", a, a, b, b, d, d)

	// the third return value of Point3D(a,b,c) is assigned to
	// parameter zz in function Cc
	cp := Cc(Point3D(a, b, c))
	fmt.Printf("%v*%v + %v*%v + %v*%v = %v*%v\n", a, a, b, b, c, c, cp, cp)

	// more return values than function parameters
	f := func(a, b, c, d float64) (float64, float64, float64, float64) {
		return a, b, c, d
	}

	// when f returns, `5, 6` are assigned to zz in Cc
	cf := Cc(f(3, 4, 5, 6)) // equivalent to: Cc(3, 4, 5, 6)
	fmt.Printf("%v\n", cf)
}
