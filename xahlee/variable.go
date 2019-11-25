package main

import "fmt"

// http://xahlee.info/golang/golang_variable.html
func main() {
	var name string
	var age int
	var height float32
	var balance float64
	var married bool
	var titles []string

	var address, city string // declare multiple variables of one type

	// in golang, default values have another name: zero value
	// the zero values are:
	//   0 => for numeric types
	//   "" => for string type
	//   false => for bool type
	//   nil => for pointers, functions, interfaces, slices, channels, and maps
	fmt.Printf("%#v\n", name)			// ""
	fmt.Printf("%#v\n", age)			// 0
	fmt.Printf("%#v\n", height)		// 0
	fmt.Printf("%#v\n", balance)	// 0
	fmt.Printf("%#v\n", married)	// false
	fmt.Printf("%#v\n", titles)		// []string(nil)

	fmt.Printf("%#v, %#v\n", address, city) // "", "", ""
}

