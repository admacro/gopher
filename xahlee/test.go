// Go programs start running in package named "name", with function "main"
package main

import "fmt"
import "math"

// line comment start with two slashes
/* block comment
alternative syntax for import
import ("fmt"; "math"; "etc")
*/

// duplicate imports are not allowed:
//   fmt redeclared as imported package name

func main() {
	fmt.Println("i love cats");
	fmt.Println(3 + 4) // semicolon is optional when is at the end of a line

	// exported (public) names, begin with upper case
	// any unexported (private) names are not accessible from outside the package
	fmt.Printf("%v\n", math.Sqrt(9))
	fmt.Printf("Pi = %v\n", math.Pi)

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
	fmt.Printf("value => type: %#v => %T, %#v => %T\n", name, name, age, age)
	fmt.Printf("80%% of the software in the world is garbage.\n")
}

