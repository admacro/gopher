// https://golang.org/ref/spec#Composite_literals
// https://golang.org/ref/spec#Type_assertions
package main

import (
	"fmt"
	"reflect"
)

func main() {
	// array literal
	// An element without a key uses the previous element's index plus one.
	// If the first element has no key, its index is zero
	a := [3]int{1, 2, 3}
	fmt.Println(a)

	// An element with a key uses the key as its index
	ak := [3]int{1: 111, 2: 222}
	fmt.Println(ak)

	// a key can be any value that's representable by a value of type int, e.g. rune, hex
	// however, the length of the array must cover the maximum of all keys
	ar := [128]int{'a': 111, '\x65': 222}
	ai := [128]int{97: 111, 101: 222} // 'a' is equivalent to 97, '\x65' to 101
	fmt.Println(ar == ai)

	// The notation ... specifies an array length equal
	// to the maximum element index plus one.
	ad := [...]int{1, 8, 9, 39, 94, 222, 45}
	fmt.Println(ad)

	// slice literal
	sl := []int{1, 8, 9}
	fmt.Printf("value: %#v, len: %d, cap: %d\n", sl, len(sl), cap(sl))

	// literals that contain elements of pointer type
	type Point struct{ x, y int }
	pa := [3]*Point{{1, 3}, {}} // same as [3]*Point{&Point{1,3}, &Point{}}
	ma := map[string]*Point{"start": {0, 0}, "end": {23, 45}}
	fmt.Printf("value: %#v, len: %d, cap: %d\n", pa, len(pa), cap(pa))
	fmt.Printf("value: %#v, len: %d\n", ma, len(ma))

	// parsing ambiguity (No more as of 2020apr1)
	// no need to enclose literal with parenthesis: ([]int{1, 2}[1]) or (2 == []int{1, 2}[1])
	if 2 == []int{1, 2}[1] {
		fmt.Println("2 == []int{1, 2}[1]")
	}

	// slice and map initialization
	// the zero value for a slice or map type is nil
	// a slice or map must be initialized to hold elements, which means storage
	// must be allocated beforehand
	// this is similar with ArrayList in Java:
	//     ArrayList<String> list // declared (value: nil), but storage not allocated
	//     ArrayList<String> list = new ArrayList<String>() // declared with storage allocation

	// declared but uninitialized (storage not allocated)
	var s []string

	// declared but uninitialized (storage not allocated, compare with new in Java)
	// ss is a pointer points to a slice with value nil
	ss := new([]string)

	// declared and initialized (storage allocated, but still an empty slice without any elements)
	sss := []string{}

	// declared and initialized (storage allocated with 5 string elements set to zero value of string type)
	ssss := make([]string, 5)

	printInfo(s, *ss, sss, ssss)

	// the same applies to map as well
	var p map[int]string
	pp := new(map[int]string)
	ppp := map[int]string{}
	pppp := make(map[int]string, 5)
	printInfo(p, *pp, ppp, pppp)
}

// for getting type from interface see
// https://stackoverflow.com/questions/20170275/how-to-find-a-type-of-an-object-in-go
func printInfo(intrs ...interface{}) {
	for _, intr := range intrs {
		val := reflect.ValueOf(intr)
		switch intr.(type) {
		case []string:
			{
				// Type Assertion
				// https://golang.org/ref/spec#Type_assertions
				//
				// x.(T) is called type assertion
				// The expression asserts that the dynamic type of x is T, which means T
				// must implement the (interface) type of x
				// The value of the expression is the value stored in x and its type is T
				// In other words, the type of the expression x.(T) is type T in a correct program
				s := intr.([]string)
				fmt.Printf("value: %#v, len: %d, cap: %d\n", val, len(s), cap(s))
			}
		case map[int]string:
			// v, ok := x.(T)
			// ok is true if the type assertion holds, otherwise if false
			// the value of v is the zero value of type T
			m, ok := intr.(map[int]string)
			if ok {
				// cap() does not support map
				fmt.Printf("value: %#v, len: %d\n", val, len(m))
			}
		}
	}
}
