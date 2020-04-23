// https://golang.org/ref/spec#Size_and_alignment_guarantees
package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type St struct {
	i int
	b bool
	f float32
	s string
	r rune
}

func main() {
	// alignment of different types
	var (
		// bool
		bl bool

		// integer
		a int
		b int8
		c int16
		d int32
		e int64
		f uint8
		g uint16
		h uint32
		i uint64

		// floating-point
		f1 float32
		f2 float64

		// complex
		c1 complex64
		c2 complex128

		// byte, rune, string
		by  byte // uint8
		rn  rune // int32
		str string
	)

	fmt.Printf("%10T [size: %2v, alignment: %v]\n", bl, unsafe.Sizeof(bl), unsafe.Alignof(bl))

	fmt.Printf("%10T [size: %2v, alignment: %v]\n", a, unsafe.Sizeof(a), unsafe.Alignof(a))
	fmt.Printf("%10T [size: %2v, alignment: %v]\n", b, unsafe.Sizeof(b), unsafe.Alignof(b))
	fmt.Printf("%10T [size: %2v, alignment: %v]\n", c, unsafe.Sizeof(c), unsafe.Alignof(c))
	fmt.Printf("%10T [size: %2v, alignment: %v]\n", d, unsafe.Sizeof(d), unsafe.Alignof(d))
	fmt.Printf("%10T [size: %2v, alignment: %v]\n", e, unsafe.Sizeof(e), unsafe.Alignof(e))
	fmt.Printf("%10T [size: %2v, alignment: %v]\n", f, unsafe.Sizeof(f), unsafe.Alignof(f))
	fmt.Printf("%10T [size: %2v, alignment: %v]\n", g, unsafe.Sizeof(g), unsafe.Alignof(g))
	fmt.Printf("%10T [size: %2v, alignment: %v]\n", h, unsafe.Sizeof(h), unsafe.Alignof(h))
	fmt.Printf("%10T [size: %2v, alignment: %v]\n", i, unsafe.Sizeof(i), unsafe.Alignof(i))

	fmt.Printf("%10T [size: %2v, alignment: %v]\n", f1, unsafe.Sizeof(f1), unsafe.Alignof(f1))
	fmt.Printf("%10T [size: %2v, alignment: %v]\n", f2, unsafe.Sizeof(f2), unsafe.Alignof(f2))

	fmt.Printf("%10T [size: %2v, alignment: %v]\n", c1, unsafe.Sizeof(c1), unsafe.Alignof(c1))
	fmt.Printf("%10T [size: %2v, alignment: %v]\n", c2, unsafe.Sizeof(c2), unsafe.Alignof(c2))

	fmt.Printf("%10T [size: %2v, alignment: %v]\n", by, unsafe.Sizeof(by), unsafe.Alignof(by))
	fmt.Printf("%10T [size: %2v, alignment: %v]\n", rn, unsafe.Sizeof(rn), unsafe.Alignof(rn))
	fmt.Printf("%10T [size: %2v, alignment: %v]\n", str, unsafe.Sizeof(str), unsafe.Alignof(str))

	fmt.Println("----------------------------------------")

	// A struct or array type has size zero if it contains no fields (or elements,
	// respectively) that have a size greater than zero.
	type EmptySt struct{}

	var est EmptySt
	// size 0, but alignment is at least 1
	fmt.Printf("%16T [size: %v, alignment: %v]\n", est, unsafe.Sizeof(est), unsafe.Alignof(est))

	var aryEst [5]EmptySt                                                                                 // array
	fmt.Printf("%16T [size: %v, alignment: %v]\n", aryEst, unsafe.Sizeof(aryEst), unsafe.Alignof(aryEst)) // size 0

	var slcEst []EmptySt                                                                                  // slice
	fmt.Printf("%16T [size: %v, alignment: %v]\n", slcEst, unsafe.Sizeof(slcEst), unsafe.Alignof(slcEst)) // size 24

	var ary [1]int
	fmt.Printf("%16T [size: %v, alignment: %v]\n", ary, unsafe.Sizeof(ary), unsafe.Alignof(ary))

	// two distinct zero-size variables may have the same address in memory
	estAddr := uintptr(unsafe.Pointer(&est))
	aryEstAddr := uintptr(unsafe.Pointer(&aryEst))
	fmt.Printf("%T<%v> == %T<%v> is %v\n", est, est, aryEst, aryEst, estAddr == aryEstAddr) //  true

	var aryEst2 [10]EmptySt
	aryEst2Addr := uintptr(unsafe.Pointer(&aryEst2))
	fmt.Printf("%T<%v> == %T<%v> is %v\n", aryEst2, aryEst2, aryEst, aryEst, aryEst2Addr == aryEstAddr) // false

	fmt.Println("----------------------------------------")

	// alignment of struct
	// For a variable x of struct type: unsafe.Alignof(x) is the largest of
	// all the values unsafe.Alignof(x.f) for each field f of x, but at least 1
	st := St{1, true, 2.3, "golang", 'c'}
	fmt.Printf("Struct %#v\n", st)
	fmt.Printf("unsafe.Alignof(%T<%#v>) = %d\n", st, st, unsafe.Alignof(st))

	// using reflect to iterate struct fields
	// see more reflect at https://pkg.go.dev/reflect
	stVal := reflect.ValueOf(st)
	for i := 0; i < stVal.NumField(); i++ {
		switch fVal := stVal.Field(i); fVal.Kind() {
		case reflect.Int:
			val := fVal.Int()                                                           // val's type is int64
			fmt.Printf("unsafe.Alignof(%T<%#v>) = %d\n", val, val, unsafe.Alignof(val)) // 8
		case reflect.Float32:
			val := fVal.Float()                                                         // val's type is float64
			fmt.Printf("unsafe.Alignof(%T<%#v>) = %d\n", val, val, unsafe.Alignof(val)) // 8
		case reflect.Bool:
			val := fVal.Bool()
			fmt.Printf("unsafe.Alignof(%T<%#v>) = %d\n", val, val, unsafe.Alignof(val)) // 1
		case reflect.String:
			val := fVal.String()
			fmt.Printf("unsafe.Alignof(%T<%#v>) = %d\n", val, val, unsafe.Alignof(val)) // 8
		case reflect.Int32:
			val := fVal.Int()                                                           // val's type is int64
			fmt.Printf("unsafe.Alignof(%T<%#v>) = %d\n", val, val, unsafe.Alignof(val)) // 8
		}
	}
}
