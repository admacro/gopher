// https://golang.org/ref/spec#Package_unsafe
// https://pkg.go.dev/unsafe
//
// package unsafe
// type ArbitraryType int  // shorthand for an arbitrary Go type; it is not a real type
// type Pointer *ArbitraryType
// func Alignof(variable ArbitraryType) uintptr
// func Offsetof(selector ArbitraryType) uintptr
// func Sizeof(variable ArbitraryType) uintptr
//
// Pointer and uintptr
// Pointer represents a pointer to an arbitrary type. There are four special
// operations available for type Pointer that are not available for other
// types:
// - A pointer value of any type can be converted to a Pointer.
// - A Pointer can be converted to a pointer value of any type.
// - A uintptr can be converted to a Pointer.
// - A Pointer can be converted to a uintptr.
// Pointer therefore allows a program to defeat the type system and read
// and write arbitrary memory. It should be used with extreme care.
package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

func main() {
	// *T1 => Pointer => *T2
	// this conversion allows reinterpreting data of one type as data of another type
	var i16 int16 = 9999
	iPtr := unsafe.Pointer(&i16)
	var i32Ptr *int32 = (*int32)(iPtr)
	var i32 int32 = *i32Ptr
	fmt.Printf("unsafe.Sizeof(%T) = %v\n", i16, unsafe.Sizeof(i16))
	fmt.Printf("unsafe.Sizeof(%T) = %v\n", i32, unsafe.Sizeof(i32))
	fmt.Printf("%T<%v> => %T<%v> => %T<%v>\n", i16, i16, iPtr, iPtr, i32, i32)

	// pointer => Pointer
	x, y, z := 1, 3.14, "go"
	xp, yp, zp := &x, &y, &z
	xPtr, yPtr, zPtr := unsafe.Pointer(xp), unsafe.Pointer(yp), unsafe.Pointer(zp)
	fmt.Printf("unsafe.Pointer(%T<%v>) = %v (%d)\n", xp, xp, xPtr, xPtr)
	fmt.Printf("unsafe.Pointer(%T<%v>) = %v (%d)\n", yp, yp, yPtr, yPtr)
	fmt.Printf("unsafe.Pointer(%T<%v>) = %v (%d)\n", zp, zp, zPtr, zPtr)

	// Pointer => pointer
	zpp := (*string)(zPtr)
	fmt.Printf("%T(%v) = %v (%d) [%v]\n", zp, zPtr, zpp, zpp, *zpp)

	// pointer to uintptr
	zpInt, err := strconv.ParseInt(fmt.Sprintf("%d", zp), 10, 64)
	if err != nil {
		panic("error parsing int")
	}
	zpUint := uintptr(zpInt)
	fmt.Printf("uintptr(%T(%v)) = %d\n", zpInt, zpInt, zpUint)

	// uintptr => Pointer (invalid usage pattern of Pointer)
	// Conversion of a uintptr back to Pointer is not valid in general.
	//
	// flycheck error: unsafeptr: possible misuse of unsafe.Pointer
	//
	// Running "go vet" can help find uses of Pointer that do not conform
	// to valid patterns, but silence from "go vet" is not a guarantee that the
	// code is valid.
	// see more at https://pkg.go.dev/unsafe?tab=doc#Pointer
	//
	// james@~$ go vet /Users/james/prog/allez/oop/unsafe.go
	// # command-line-arguments
	// prog/allez/oop/unsafe.go:106:15: possible misuse of unsafe.Pointer
	zpUintPtr := unsafe.Pointer(zpUint)
	fmt.Printf("unsafe.Pointer(%T(%v)) = %d [%v]\n", zpUint, zpUint, zpUintPtr, *(*string)(zpUintPtr))

	// Pointer => uintptr
	zPtrUint := uintptr(zPtr)
	fmt.Printf("uintptr(unsafe.Pointer(%T<%v>)) = %v (%d)\n", zp, zp, zPtrUint, zPtrUint)

	// Pointer => uintptr => Pointer (with arithmetic)
	// struct
	type St struct {
		s string
		t float32
	}
	st := St{"pi", 3.14}
	piPtr := unsafe.Pointer(&st)
	piPtrS := unsafe.Pointer(uintptr(piPtr) + unsafe.Offsetof(st.s)) // equivalent to unsafe.Pointer(&st.s)
	piPtrT := unsafe.Pointer(uintptr(piPtr) + unsafe.Offsetof(st.t)) // equivalent to unsafe.Pointer(&st.t)
	fmt.Printf("unsafe.Pointer(uintptr(piPtr) + unsafe.Offsetof(st.s)) = %v (%q)\n", piPtrS, *(*string)(piPtrS))
	fmt.Printf("unsafe.Pointer(uintptr(piPtr) + unsafe.Offsetof(st.t)) = %v (%v)\n", piPtrT, *(*float32)(piPtrT))
	// array
	a := [3]int{1, 2, 3}
	aPtr := unsafe.Pointer(&a)
	for i := range a {
		v := unsafe.Pointer(uintptr(aPtr) + uintptr(i)*unsafe.Sizeof(&a[0]))
		fmt.Printf("a[%v] = %v\n", i, *(*int)(v))
	}
}
