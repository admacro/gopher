package main

import "fmt"

func main() {
	// variable declaration, without type (see variable.go for more)
	var name = "James"
	var age = 35

	// use fmt package for printing
	// %v -> any value (print in human readable form)
	fmt.Printf("My name is %v, and I'm %v years old.\n", name, age)
	// %#v -> print in golang syntax (e.g. string value in quotes)
	fmt.Printf("variable name is %#v, age is %#v\n", name, age)

	// other fmt placeholders are:
	// %+v also print struct field if value is a struct
	//   + seems to mean extra/plus :D (see struct.go for details)
	// %T print type of the value (T means type, obviously)
	// %% print % (escaping)
	fmt.Printf("value => type: %#v => %T, %#v => %T, %#v => %T\n", name, name, age, age, 3.14, 3.14)
	fmt.Printf("80%% of the software in the world is garbage.\n")

	// format something to string
	name = "book"
	// returns the formatted string instead of printing
	var msg = fmt.Sprintf("I'm a %v to be returned.", name)
	fmt.Println(msg)
	fmt.Println(msg == "I'm a book to be returned.") // true
}

