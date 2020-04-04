// https://golang.org/ref/spec#Address_operators
// For an operand x of type T, the address operation &x generates a pointer
// of type *T to x.
package main

import "fmt"

func main() {
	// addressable operands
	// variable
	var s = "hello"
	sp := &s
	fmt.Println(sp)

	// pointer indirection
	// *sp is a point indirection
	fmt.Println(&(*sp)) // same as above

	// slice indexing operation
	slice := []string{"bonjour", "madame"}
	fmt.Println(&slice[1])
	fmt.Println(&(slice[1])) // same as above
	// fmt.Println((&slice)[1]) // invalid operation: (&slice)[1] (type *[]string does not support indexing)

	// field selector of an addressable struct operand
	type St struct{ s string }
	st := St{"merci"}
	fmt.Println(&st.s)
	fmt.Println(&(st.s)) // same as above
	fmt.Println((&st).s) // fields of a struct s can be accessecd by its pointer *s

	// array indexing operation of an addressable array
	a := [...]St{st, st, st}
	fmt.Printf("%T: %v\n", &a[1], &a[1]) // *main.St
	fmt.Println(&(a[1]))                 // same as above

	// composite literal

}
