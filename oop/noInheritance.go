// https://golang.org/doc/effective_go.html#embedding
package main

import "fmt"

type Animal struct {
	name string
}

func NewAnimal() *Animal {
	return &Animal{name: "Animal"}
}

func (a *Animal) Walk() {
	fmt.Printf("%s walking...\n", a.name)
}

// noraml composition
// you need to provide forwarding method to access the underlying method in Animal
type Cat struct {
	animal Animal
}

func (c *Cat) Walk() {
	fmt.Print("calling animal to walk... ")
	c.animal.Walk()
}

// embedding
//
type Dog struct {
	Animal
}

func (a *Dog) Walk() {
	fmt.Printf("%s running...\n", a.name)
}

func main() {
	a := NewAnimal()
	a.Walk()

	c := Cat{animal: Animal{name: "Cat"}}
	c.Walk() // Without forwarding method: c.Walk undefined (type *Cat has no field or method Walk)
	c.animal.Walk()

	// the field name of an embededded type is same as the name of the type
	// but if you put `Animal Animal` in the `Dog` struct, that'll have the
	// same effect as a composition `a Animal`
	d := Dog{Animal: Animal{name: "Dog"}}
	d.Walk()
}
