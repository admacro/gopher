// https://golang.org/ref/spec#Package_unsafe
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	str := "hello Golang"
	fmt.Println("Results:", unsafe.Sizeof(str))
	fmt.Println("Results:", len(str))

	// Computer architectures may require memory addresses to be aligned;
	// that is, for addresses of a variable to be a multiple of a factor,
	// the variable's type's alignment.
	// The function Alignof takes an expression denoting a variable of any
	// type and returns the alignment of the (type of the) variable in bytes.
	// see more at http://en.wikipedia.org/wiki/Data_structure_alignment
	x := 1
	fmt.Println("Results:", unsafe.Alignof(x))

	type St struct {
		f int
		g float32
	}
	s := St{1, 2.3}
	fmt.Println("Results:", unsafe.Offsetof(s.f))
	fmt.Println("Results:", unsafe.Offsetof(s.g))
}
