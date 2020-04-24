// https://golang.org/doc/effective_go.html#type_switch
// https://golang.org/ref/spec#Switch_statements
// https://golang.org/ref/spec#Type_assertions
// https://golang.org/ref/spec#Package_unsafe
package main

import (
	"fmt"
)

func printInfo(intrs ...interface{}) {
	for _, i := range intrs {
		switch v := i.(type) {
		default:
			// inside the default clause, type of v is type of i (interface{})
			// thus, direct conversion of v to any other type, except interface{},
			// is illegal
			// to extract the actual type and value from v, type assertion is needed,
			// as told by the compiler
			// normally, you should not do type assertion manually in the default clause
			// of a type switch since that's what type switch does for you automatically
			//
			// error: cannot convert v (type interface {}) to type float32: need type assertion
			// vv := float32(v)
			if vv, ok := i.(float32); ok {
				fmt.Printf("default: %T %v (type assertion)\n", vv, vv)
			}

			// however, when printing, fmt.Printf does type assertion automatically
			// therefore, what's printed is the dynamic type (and value) stored in v
			// see source code of pp.printArg in fmt package for more details
			fmt.Printf("default: %T %v\n", v, v)

		case int8:
			// inside a case clause, v is same as vv where vv is
			// the result of type assertion: vv, ok := i.(int8)
			// thus, type of v is type of vv which is int8
			// and v could be converted to other integer types
			fmt.Printf("case: %T %v (as int: %v)\n", v, v, int(v))
		}
	}
}

func main() {
	var a int8 = 1
	var b float32 = 3.14
	printInfo(a, b)

	type St struct{ s string }
	printInfo([]St{St{"go"}})
}
