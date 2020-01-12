package main

import "fmt"

// struct is similar to struct in c and ruby, map in golang
// it's a user-defined type
// you define a collection of fields, each consists of a name and a type
// once it's defined, the struct definition is done and the fileds are fixed
// you can create new struct with values from the struct definition

// struct is very important in golang, because oop's concept of interface and
// methods are based on it
func main() {
	// define a struct
	// field names must be unique
	type Person struct {
		name string
		age  int
	}

	type Position struct {
		x, y int
	}

	// create a strut
	var p = Person{"Joker", 40} // if field names are omitted, all fields must be present in the order defined
	// var p = Person{"Joker"}				// err: too few values in Person literal
	var pp = Person{} // this is ok

	fmt.Println(p)          // {Joker 40}
	fmt.Printf("%+v\n", p)  // {name:Joker age:40}
	fmt.Printf("%#v\n", p)  // main.Person{name:"Joker", age:40}
	fmt.Printf("%#v\n", pp) // main.Person{name:"", age:0}

	var bb = Person{name: "baby"}
	fmt.Println(bb)                           // {baby 0} ommited field will have zero value of its type
	bb.age = 3                                // modify/update/fill struct field
	fmt.Printf("{%v, %v}\n", bb.name, bb.age) // {baby, 3}
}
