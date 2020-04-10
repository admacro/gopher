// https://golang.org/ref/spec#Switch_statements
// https://golang.org/ref/spec#Type_assertions
// see compositeliteral.go#printInfo for another example
package main

import (
	"fmt"
)

func main() {
	var (
		// rune, string
		c = 'G'
		s = "golang"

		// integer
		i       = 1 // dynamic type is int
		j int   = 1
		k int64 = 1

		// floating-point
		f float32 = 3.1415926
		g float64 = 3.1415926
	)
	printInfo(c, s)
	printInfo(i, j, k)
	printInfo(f, g)
	printInfo(nil)
	printInfo(true)
}

func printInfo(intrs ...interface{}) {
	for _, intr := range intrs {
		// v is the value intr holds
		// the variable is declared at the end of the TypeSwitchCase
		// in the implicit block of each clause
		switch v := intr.(type) {
		case rune:
			// here, vv is same as v above
			vv := intr.(rune)
			fmt.Printf("%T(%c)\n", v, v) // will print int32 as its type, not rune
			fmt.Printf("%v == %v is %v\n", v, vv, v == vv)
		case string:
			// ok is true if the type assertion holds which means intr is of type string,
			// otherwise it's false and the value of vv is the zero value of type string
			vv, ok := intr.(string)
			fmt.Printf("%T(%s) (ok = %v)\n", v, vv, ok)
		case int, int8, int16, int64: // int32 is duplicate with rune above
			fmt.Printf("%T(%d)\n", v, v)
		case float32, float64:
			fmt.Printf("%T(%3.2f)\n", v, v)
		case nil: // interface values could be nil
			fmt.Printf("%T(%v)\n", v, v)
		default:
			fmt.Printf("%T(%v) (Type %T is not supported, fallback to fmt.Printf)\n", v, v, v)
		}
	}
}
