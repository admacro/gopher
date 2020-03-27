// https://golang.org/ref/spec#Map_types
// The comparison operators == and != must be fully defined for operands of
// the key type; thus the key type must not be a function, map, or slice.
// If the key type is an interface type, these comparison operators must be
// defined for the dynamic key values; failure will cause a run-time panic.
package main

import "fmt"

type Id interface {
	String() string
}
type Order struct {
	id string
}

func (order *Order) String() string {
	return fmt.Sprintf("OrderId#%v", order.id)
}

func main() {
	orderStatus := make(map[Id]string)
	var id Id = &(Order{"12345"})
	orderStatus[id] = "shipped"
	fmt.Println(orderStatus)
}
