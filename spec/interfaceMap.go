// https://golang.org/ref/spec#Map_types
// The comparison operators == and != must be fully defined for operands of
// the key type; thus the key type must not be a function, map, or slice.
// If the key type is an interface type, these comparison operators must be
// defined for the dynamic key values; failure will cause a run-time panic.

// https://golang.org/ref/spec#Comparison_operators
// Interface values are comparable. Two interface values are equal if they
// have identical dynamic types and equal dynamic values or if both have
// value nil.
// Struct values are comparable if all their fields are comparable. Two
// struct values are equal if their corresponding non-blank fields are equal.
package main

import "fmt"

type Id interface {
	IdStr() string
}

type Order struct {
	id string
}

func (order Order) IdStr() string {
	return fmt.Sprintf("OrderId#%s", order.id)
}

// this is for demo purpose
// for singleton implementation in Go, see
// https://stackoverflow.com/questions/1823286/singleton-in-go
type Singleton struct {
	id string
}

func (s *Singleton) IdStr() string {
	return fmt.Sprintf("SingeltonId(%v)#%s", s, s.id)
}

func testValue() {
	// map key type is interface type Id
	// keys are compared by comparing their dynamic values
	orderStatus := make(map[Id]string)

	// dynamic key type is Order, dynamic values are Order values
	// Order values are compared by comparing their fields values
	var id Id = Order{"12345"}
	orderStatus[id] = "ordered"
	fmt.Printf("%#v\n", orderStatus)

	// distinct Order value with same field value, thus same key
	var idCopy Id = Order{"12345"}
	orderStatus[idCopy] = "shipped"
	fmt.Printf("%#v\n", orderStatus)

	// same as above
	orderStatus[Order{"12345"}] = "delivered"
	fmt.Printf("%#v\n", orderStatus)
}

func testPointer() {
	// map key type is interface type Id
	// keys are compared by comparing their dynamic values
	singletonStates := make(map[Id]string)

	// dynamic key type is *Singleton, dynamic key values are *Singleton values
	// *Singleton values are addresses (hexdecimal), they are compared directly
	// field values of Singleton values are not compared
	var id Id = &Singleton{"67890"}
	singletonStates[id] = "OK"
	fmt.Printf("%#v\n", singletonStates)

	// distinct Singleton value with same field value
	// thus, address is different with above
	id2 := &Singleton{"67890"}
	var idCopy Id = id2
	singletonStates[idCopy] = "OK"
	fmt.Printf("%#v\n", singletonStates)
	singletonStates[id2] = "OKOK"
	fmt.Printf("%#v\n", singletonStates)

	// keys are equal if the addresses are equal
	// singletonStates[&(*idCopy)] = "OK Copy" // error: invalid indirect of idCopy (type Id)
	singletonStates[&(*id2)] = "OK Copy"
	fmt.Printf("%#v\n", singletonStates)
}

func main() {
	testValue()
	testPointer()
}
