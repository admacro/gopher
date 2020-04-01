// https://golang.org/ref/spec#Composite_literals
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
	// ss points to a slice with value nil
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
func printInfo(ss ...interface{}) {
	for _, s := range ss {
		sv := reflect.ValueOf(s)
		switch s.(type) {
		case []string:
			{
				sl := s.([]string)
				fmt.Printf("value: %#v, len: %d, cap: %d\n", sv, len(sl), cap(sl))
			}
		case map[int]string:
			// cap() does not support map
			fmt.Printf("value: %#v, len: %d\n", sv, len(s.(map[int]string)))
		}
	}
}
