package main

import "fmt"

func main() {
	// [n]T is a type in Go
	// it's an array of n values of type T
	var names [3]string						// length is 3
	names[0] = "James"

	// an array's length (n) is part of its type, so arrays cannot be resized
	fmt.Printf("%T\n", names)			// [3]string

	fmt.Printf("%d\n", len(names))
	fmt.Printf("%d\n", cap(names)) // same as len for array

	for i, name := range names {
		fmt.Printf("names[%d] = %#v\n", i, name)
	}

	// array literal and loop
	var cities = [3]string{"Shanghai", "San Francisco", "Wellington"}

	for i, city := range cities {
		fmt.Printf("city[%d] = %#v\n", i, city)
	}

	// array[:] returns a slice of array
	var slice = cities[:]
	fmt.Printf("%T\n", slice)			 // []string
	fmt.Printf("%d\n", len(slice)) // 3
}
