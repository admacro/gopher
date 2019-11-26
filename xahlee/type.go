package main

import "fmt"
import "reflect"

func main() {
	var (
		// int, uint, and uintptr are usually 32bit wide on
		// 32bit systems, and 64bit wide on 64bit systems

		a int = 1
		b int8 = 2
		c int16 = 3
		d int32 = 4									// alias rune
		e int64 = 5

		ua uint = 1
		ub uint8 = 2								// alias byte
		uc uint16 = 3
		ud uint32 = 4
		ue uint64 = 5

		fa float32 = 3.14
		fb float64 = 0.618

		ca complex64 = 1.22233
		cb complex128 = 1.22233

		ptr uintptr
	)

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

	fmt.Printf("%T\n", fa)
	fmt.Printf("%T\n", fb)

	fmt.Printf("%v\n", ca)				// (1.22233+0i)
	fmt.Printf("%T\n", cb)
	
	fmt.Printf("%T\n", ptr)

	// type conversion
	// newtype(v) converts v to newtype
	fmt.Printf("%T\n", int32(ua))
	fmt.Printf("%T\n", float64(fa))

	// another way to get the type of a variable is to use
	// reflect.TypeOf(v)
	var ca64 = complex64(ca)
	fmt.Println(reflect.TypeOf(ca64))
}
