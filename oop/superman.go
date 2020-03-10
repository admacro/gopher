// i don't know what the hell this is LOL
package main

import "fmt"

type Human struct {
	Gender string
	Age    int
}

type Introducer interface {
	Introduce()
}

func (h *Human) Introduce() {
	fmt.Printf("I'm %s. I'm %d years old.\n", h.Gender, h.Age)
}

type Man Human

// overrides Human#Introduce()
func (m *Man) Introduce() {
	fmt.Printf("I'm a man. I'm %d years old.\n", m.Age)
}

type Superman Man

func (s *Superman) Fly() {
	fmt.Println("I can fly.")
}

func main() {
	h := Human{Gender: "male", Age: 25}
	h.Introduce() // I'm male. I'm 25 years old.

	man := Man{Gender: "male", Age: 25}
	man.Introduce() // I'm a man. I'm 25 years old.

	// To call method of parent type, you need to convert to parent type first
	// This explicit type conversion makes it easier to know the source of the
	// method being called just by reading the source code
	// In Java, you use super.someMethod() or super.super.someMethod() which is
	// not very helpful
	superman := Superman{Gender: "male", Age: 35}

	normalMan := Man(superman)
	normalMan.Introduce() // I'm a man. I'm 35 years old.

	normalHuman := Human(superman)
	normalHuman.Introduce() // I'm male. I'm 35 years old.

	superman.Fly() // I can fly.
}
