package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// Person implements the String() method in Stringer type
// The fmt package (and many others) look for the interface to print values
// It's similar to toString() in Java
func (p Person) String() string {
	return fmt.Sprintf("Name=%v, Age=%d", p.Name, p.Age)
}

func main() {
	var james = Person{"James", 34}
	var parry = Person{"Parry", 24}
	fmt.Println(james)
	fmt.Println(parry)
}
