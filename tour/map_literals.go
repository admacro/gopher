package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

// initialize a variable of type map[string]Vertex
var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["shanghai"] = Vertex{31.25516, 121.4747}
	fmt.Println(m)

	// literals
	var n = map[string]Vertex{
		"wellington": {
			41.2925488, 174.7733654,
		},
		// If the top-level type is just a type name, you
		// can omit it from the elements of the literal.
		"san francisco": {
			37.779379, 122.418433,
		},
	}
	fmt.Println(n)
}
