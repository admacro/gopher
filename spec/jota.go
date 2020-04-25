// https://golang.org/ref/spec#Function_literals
// A function literal represents an anonymous function
// Function literals are closures: they may refer to variables defined
// in a surrounding function. Those variables are then shared between the
// surrounding function and the function literal, and they survive as long
// as they are accessible.
// see ../xahlee/closure.go for more on closure
// see ./kota.go for a more versatile version of *ota (LOL)
package main

import "fmt"

// this is a function type
type Jota func() int

func makeJota() Jota {
	var v int = -1

	// assgin a function literal to f
	f := func() int {
		v = v + 1
		return v
	}
	return f
}

type Kota func(...bool) int

func makeKota(init ...int) Kota {
	var v int = -1

	length := len(init)
	if length == 1 {
		v = init[0] - 1
	} else if length > 1 {
		panic("makeKota can only have exactly one value if provided")
	}

	// assgin a function literal to f
	f := func(pause ...bool) int {
		var p bool
		length := len(pause)
		if length == 1 {
			p = pause[0]
		} else if length > 1 {
			panic("kota can only have exactly one parameter if provided")
		}
		if !p {
			v = v + 1
		}
		return v
	}
	return f
}

func main() {
	jota := makeJota()
	fmt.Println(jota())
	fmt.Println(jota())
	fmt.Println(jota())

	// unlike iota
	// multiple uses of jota in the same statement all have different values
	// the value is increamented every call with jota()
	a, b := jota(), jota()
	fmt.Println(a, b) // 3, 4

	// Kota
	kota := makeKota()
	fmt.Println(kota())
	fmt.Println(kota())
	fmt.Println(kota())

	// custom init
	kota = makeKota(5) // start from 5
	fmt.Println(kota())
	fmt.Println(kota())
	fmt.Println(kota()) // 7

	// increment control
	a, b = kota(true), kota(true)
	fmt.Println(a, b) // 7, 7
}
