// https://golang.org/ref/spec#Function_literals
// see ./jota.go for a simpler version
package main

import "fmt"

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
	kota := makeKota() // start from 0 by default
	fmt.Println(kota())
	fmt.Println(kota())
	fmt.Println(kota())

	// custom init
	kota = makeKota(5) // make a new kota that starts from 5
	fmt.Println(kota())
	fmt.Println(kota())
	fmt.Println(kota()) // 7

	// increment control
	a, b := kota(true), kota(true)
	fmt.Println(a, b) // 7, 7
}
