// https://golang.org/ref/spec#Constant_declarations
package main

import "fmt"

const (
	// Within a parenthesized const declaration list the expression list may be omitted from any but the first one.
	// Omitting the list of expressions is equivalent to repeating the previous list.
	constant1 = "first"
	constant2 // this is equivalent to: constant2 = "first"

	constant3, constant4 = 111, 222
	constant5, constant6 // this is equivalent to: constant5, constant6 = 111, 222
)

// With the iota constant generator, this mechanism permits
// light-weight declaration of sequential values
// the value of iota depends on its position in the constants list
const (
	Lundi    = iota // iota starts from 0 if it's the first in the constant list
	Mardi           // 1 (as mentioned above, this is same as: Mardi = iota, the same applies to subsequent declarations)
	Mercredi        // 2
	Jeudi           // 3
	Vendredi        // 4
	Samedi          // 5
	Dimanche        // 6
)

// the value of iota depends on its position in the constants list
const (
	_         = "placeholder"
	__        = "placeholder"
	Monday    = iota // iota is 2 as it's the third constant in the list
	Thusday          // 3
	Wednesday        // 4
)

func main() {
	fmt.Println(constant2)
	fmt.Println(constant5)

	fmt.Println(Lundi)
	fmt.Println(Mardi)
	fmt.Println(Jeudi)
	fmt.Println(Samedi)

	fmt.Println(Monday)
	fmt.Println(Wednesday)
}
