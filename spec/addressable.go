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
	fmt.Printf("%T: %v\n", sp, sp)

	// pointer indirection
	// *sp is a point indirection
	fmt.Printf("%T: %v\n", &(*sp), &(*sp)) // same as above

	// slice indexing operation
	slice := []string{"bonjour", "madame"}
	fmt.Printf("%T: %v\n", &slice[1], &slice[1])
	fmt.Printf("%T: %v\n", &(slice[1]), &(slice[1])) // same as above
	// fmt.Println((&slice)[1]) // invalid operation: (&slice)[1] (type *[]string does not support indexing)

	// field selector of an addressable struct operand
	type St struct{ s string }
	st := St{"merci"}
	fmt.Printf("%T: %v\n", &st, &st)
	fmt.Printf("%T: %v\n", &st.s, &st.s)
	fmt.Printf("%T: %v\n", &(st.s), &(st.s)) // same as above
	fmt.Printf("%T: %v\n", (&st).s, (&st).s) // fields of a struct s can be accessecd by its pointer *s

	// array indexing operation of an addressable array
	a := [...]int{1, 2, 3}
	fmt.Printf("%T: %v\n", &a[1], &a[1])
	fmt.Printf("%T: %v\n", &(a[1]), &(a[1])) // same as above

	// composite literal
	cl := &St{"struct literal"}
	fmt.Printf("%T: %v\n", cl, cl)
	m := &map[int]string{1: "map literal"}
	fmt.Printf("%T: %v\n", m, m)
	fmt.Printf("%T: %v\n", (*m)[1], (*m)[1])
	// fmt.Printf("%T: %v\n", &(*m)[1], &((*m)[1])) // cannot take the address of (*m)[1]
}
