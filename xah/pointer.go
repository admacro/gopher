package main

import "fmt"

func main() {
	var city = "shanghai"

	// create a pointer to a variable xyz
	// &xyz
	var p2c = &city

	// pointer in golang syntax
	fmt.Printf("pointer in golang syntax: %#v\n", p2c) // (*string)(0xc000010200)

	// type of pointer
	// if a variable's type is abc, then the type of its pointers is *abc
	fmt.Printf("type of variable %T\n", city) // string
	fmt.Printf("type of pointer: %T\n", p2c)  // *string

	// address of variable
	fmt.Printf("address of variable: %v\n", p2c) // 0xc000010200

	// get the value of the variable to which the pointer is pointing
	// *pointer (dereferencing/pointer indirection)
	fmt.Printf("*p2c = %#v\n", *p2c)                    // "shanghai"
	fmt.Printf("(*p2c == city) => %#v\n", *p2c == city) // true

	// pointer to struct
	type Position struct {
		x, y int
	}
	var pos = Position{3, 4}
	var ptr = &pos
	fmt.Printf("%T\n", pos) // main.Position
	fmt.Printf("%T\n", ptr) // &main.Position

	// if i have a struct pointer, this is how struct fields
	// are typically accessed, but it's cubersome
	(*ptr).x = 11
	fmt.Printf("%#v\n", pos) // main.Position{x:11, y:4}

	// so, Go permits struct fields to be accessed through a
	// struct pointer directly
	ptr.x = 5                // update field of associated struct
	fmt.Printf("%#v\n", pos) // main.Position{x:5, y:4}

	// why use pointer?
	// the purpose of using pointer is mostly for speeding up computation

	// when f is called, golang will make a copy of x and pass it to f
	// making a copy slows things down
	var f = func(x string) string {
		x += "!!!"
		return x
	}
	var salut = "hello"
	fmt.Println(f(salut)) // hello!!!
	fmt.Println(salut)    // hello

	// pointer as argument
	// passing pointer to a function doesn't do value copying
	// however, now you can modify the variable directly in the
	// function body by dereferencing the pointer
	var fp = func(x *string) string {
		*x += "!!!"
		return *x
	}
	fmt.Println(fp(&salut)) // hello!!!
	fmt.Println(salut)      // hello!!!
}
