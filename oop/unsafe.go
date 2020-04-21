// https://golang.org/ref/spec#Package_unsafe
//
// package unsafe
// type ArbitraryType int  // shorthand for an arbitrary Go type; it is not a real type
// type Pointer *ArbitraryType
// func Alignof(variable ArbitraryType) uintptr
// func Offsetof(selector ArbitraryType) uintptr
// func Sizeof(variable ArbitraryType) uintptr
//
// doc: https://pkg.go.dev/unsafe
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var (
		b    bool
		i    int8
		ui   uint16
		f    float32
		r    rune
		str0 string
		str1 = "hello Golang"
	)

	// the size does not include any memory possibly referenced by str
	// it returns the size of the string descriptor, not the size of
	// the memory referenced by the slice.
	// thus Sizeof returns the same size for str0 (zero value) and str1
	fmt.Printf("unsafe.Sizeof(%T<%v>) = %d\n", b, b, unsafe.Sizeof(b))          // 1
	fmt.Printf("unsafe.Sizeof(%T<%d>) = %d\n", i, i, unsafe.Sizeof(i))          // 1
	fmt.Printf("unsafe.Sizeof(%T<%d>) = %d\n", ui, ui, unsafe.Sizeof(ui))       // 2
	fmt.Printf("unsafe.Sizeof(%T<%f>) = %d\n", f, f, unsafe.Sizeof(f))          // 4
	fmt.Printf("unsafe.Sizeof(%T<%c>) = %d\n", r, r, unsafe.Sizeof(r))          // 4
	fmt.Printf("unsafe.Sizeof(%T<%q>) = %d\n", str0, str0, unsafe.Sizeof(str0)) // 16
	fmt.Printf("unsafe.Sizeof(%T<%q>) = %d\n", str1, str1, unsafe.Sizeof(str1)) // 16 (same as str0)
	fmt.Printf("len(%T<%q>) = %d\n", str1, str1, len(str1))                     // 12

	// Computer architectures may require memory addresses to be aligned;
	// that is, for addresses of a variable to be a multiple of a factor,
	// the variable's type's alignment.
	// The function Alignof takes an expression denoting a variable of any
	// type and returns the alignment of the (type of the) variable in bytes.
	// see more at http://en.wikipedia.org/wiki/Data_structure_alignment
	fmt.Printf("unsafe.Alignof(%T<%v>) = %d\n", b, b, unsafe.Alignof(b))          // 1
	fmt.Printf("unsafe.Alignof(%T<%q>) = %d\n", i, i, unsafe.Alignof(i))          // 1
	fmt.Printf("unsafe.Alignof(%T<%q>) = %d\n", ui, ui, unsafe.Alignof(ui))       // 2
	fmt.Printf("unsafe.Alignof(%T<%f>) = %d\n", f, f, unsafe.Alignof(f))          // 4
	fmt.Printf("unsafe.Alignof(%T<%c>) = %d\n", r, r, unsafe.Alignof(r))          // 4
	fmt.Printf("unsafe.Alignof(%T<%q>) = %d\n", str0, str0, unsafe.Alignof(str0)) // 8
	fmt.Printf("unsafe.Alignof(%T<%q>) = %d\n", str1, str1, unsafe.Alignof(str1)) // 8 (same as str0)

	// Offsetof returns the number of bytes between the
	// start of the struct and the start of the field
	type St struct {
		i int
		b bool
		f float32
		s string
		r rune
	}
	st := St{1, true, 2.3, "golang", 'c'}
	fmt.Printf("%#v\n", st)
	fmt.Printf("unsafe.Offsetof(%T<%q>) = %d\n", st.i, st.i, unsafe.Offsetof(st.i)) // 0 (i is the first field, thus 0)
	fmt.Printf("unsafe.Offsetof(%T<%v>) = %d\n", st.b, st.b, unsafe.Offsetof(st.b)) // 8
	fmt.Printf("unsafe.Offsetof(%T<%f>) = %d\n", st.f, st.f, unsafe.Offsetof(st.f)) // 12
	fmt.Printf("unsafe.Offsetof(%T<%q>) = %d\n", st.s, st.s, unsafe.Offsetof(st.s)) // 16
	fmt.Printf("unsafe.Offsetof(%T<%q>) = %d\n", st.r, st.r, unsafe.Offsetof(st.r)) // 32
}
