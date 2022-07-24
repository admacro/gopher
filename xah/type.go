package main

import (
	"fmt"
	"reflect"
)

func main() {
	var (
		// int, uint, and uintptr are usually 32bit wide on
		// 32bit systems, and 64bit wide on 64bit systems

		// when you need an integer value, you should use int
		// unless you have very specific reason to use a sized
		// or unsigned integer type.

		p bool = true

		// An int can store at maximum a 64-bit integer, and sometimes less.
		big int   = 1<<63 - 1
		a   int   = 1
		b   int8  = 2
		c   int16 = 3
		d   int32 = 4 // alias rune
		e   int64 = 5

		ua uint   = 1
		ub uint8  = 2 // alias byte
		uc uint16 = 3
		ud uint32 = 4
		ue uint64 = 5

		bt byte = 123 // alias for uint8
		r  rune = 96  // alias for int32

		fa float32 = 3.14
		fb float64 = 0.618

		ca complex64  = 1.22233
		cb complex128 = 1.22233

		ptr uintptr
	)

	fmt.Printf("%T\n", p)
	fmt.Printf("%T: %v\n", big, big)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", d)
	fmt.Printf("%T\n", e)

	fmt.Printf("%T\n", ua)
	fmt.Printf("%T\n", ub)
	fmt.Printf("%T\n", uc)
	fmt.Printf("%T\n", ud)
	fmt.Printf("%T\n", ue)

	fmt.Printf("%T\n", bt) // uint8
	fmt.Printf("%T\n", r)  // int32

	fmt.Printf("%T\n", fa)
	fmt.Printf("%T\n", fb)

	fmt.Printf("%v\n", ca) // (1.22233+0i)
	fmt.Printf("%T\n", cb)

	fmt.Printf("%T\n", ptr)

	// type conversion
	// T(v) converts value v to type T
	fmt.Printf("%T\n", int32(ua))
	fmt.Printf("%T\n", float64(fa))

	// another way to get the type of a variable is to use
	// reflect.TypeOf(v)
	var ca64 = complex64(ca)
	fmt.Println(reflect.TypeOf(ca64))
}
