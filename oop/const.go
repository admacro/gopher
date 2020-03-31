// https://golang.org/ref/spec#Constant_declarations
package main

import "fmt"

const (
	// Within a parenthesized const declaration list the expression list may be omitted from any but the first one.
	// Omitting the list of expressions is equivalent to repeating the previous list.
	constant1 = 123
	constant2 // this is equivalent to: constant2 = 123

	constant3, constant4 = 111, 222
	constant5, constant6 // this is equivalent to: constant5, constant6 = 111, 222
)

// Within a constant declaration, the predeclared identifier iota represents
// successive untyped integer constants. Its value is the index of the respective
// ConstSpec in that constant declaration, starting at zero.

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

const (
	empty, zero, none = iota, iota, iota // multiple uses of iota in the same ConstSpec all have the same value
	first, one, top
)

func main() {
	fmt.Println(constant2)
	fmt.Println(constant5)
	fmt.Println(Lundi, Mardi, Jeudi, (Samedi))
	fmt.Println(Monday, Wednesday)
	fmt.Println(empty, none, one, top)
}
