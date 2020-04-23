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

// noraml composition:
//   - some code resue
//   - you need to provide forwarding method (delegation) to access the underlying method in Animal
type Cat struct {
	animal Animal
}

// delegation
func (c *Cat) Walk() {
	fmt.Print("calling animal to walk... ")
	c.animal.Walk()
}

// composition and embedding: field and method reuse
type Dog struct {
	Animal
}

// interface: polymorphism and dynamic dispatch
type Sleeper interface {
	Sleep()
}

type Human struct {
	Animal
}

func (a *Animal) Sleep() {
	fmt.Printf("%s sleeping...\n", a.name)
}

func (d *Dog) Sleep() {
	fmt.Printf("%s sleeping...\n", d.name)
}

func (h *Human) Sleep() {
	fmt.Printf("%s sleeping...\n", h.name)
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

	h := Human{}
	h.name = "Human"

	// so interface can be seen as a means to group types by behaviour
	// if
	//     - A and B do x
	//     - B and C do y
	//     - interface Xer has method x
	//     - interface Yer has method y
	// then
	//     - A and B are grouped as Xers
	//     - Or, B and C are bother Yers
	//
	// When I see a bird that walks like a duck and swims like a duck and
	// quacks like a duck, I call that bird a duck.
	//                                              –James Whitcomb Riley
	sleepers := []Sleeper{&d, &h}
	for _, s := range sleepers {
		s.Sleep()
	}
}
