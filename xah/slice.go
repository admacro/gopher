// https://golang.org/ref/spec#Slice_types
package main

import (
	"fmt"
	"reflect"
	"strings"
)

// print slice length, capacity, and items line by line
func print_slice_info(s []string, title string) {
	fmt.Printf("\n#%v----------------------------\n", title)
	fmt.Printf("length: %d, capacity: %d \n", len(s), cap(s))
	for i, item := range s {
		fmt.Printf("%v[%d]: %#v\n", title, i, item)
	}
}

func main() {
	// A slice is a descriptor for a segment of contiguous elements of an underlying array.
	// It's essentially a reference to a segment of arary.
	// It's a dynamically-sized, flexible view into the elements of an array.
	// It's like array, but its length can be changed (array length is fixed).
	// In practice, slices are much more common than arrays.
	// Slice is similar to ArrayList in Java.

	// []type => slice
	// [n]type => array

	// this firstly creates a hidden array: [3]string{"Java", "Go", "Ruby"}
	// then builds a slice that references it
	var langs = []string{"Java", "Go", "Ruby"}

	fmt.Printf("Type of langs: %T\n", langs) // []string
	fmt.Printf("Langs: %#v\n", langs)
	fmt.Printf("Langs: %v\n", langs)

	langs[0] = "Javascript"
	fmt.Printf("Langs: %#v\n", langs)
	fmt.Printf("Langs: %v\n", langs)

	print_slice_info(langs, "langs") // length: 3, capacity: 3

	// To create slice, use the built-in function `make`
	// A slice created with `make` always allocates a new, hidden array to which
	// the returned slice value refers.
	// So the following two expressions are equivalent:
	//   make([]type, length, capacity) (e.g. make([]int, 5, 10))
	//   new([capacity]type)[0:length] (e.g. new([10]int)[0:5])

	// make([]type, count_n) => creates a slice of count_n items of type
	// capacity defaults to len if not specified
	var fingers = make([]string, 5)
	print_slice_info(fingers, "fingers") // length: 5, capacity: 5

	// make([]type, count_n, capacity_m) => with capacity of capacity_m items
	var followers = make([]string, 3, 5)
	print_slice_info(followers, "followers") // length: 3, capacity: 5

	// at any time, 0 <= length <= capacity
	// var followers = make([]string, 6, 5) // err: len larger than cap in make([]string)

	// slice item assignment
	for i := range followers {
		followers[i] = fmt.Sprintf("assignment #%d", i+1)
	}
	print_slice_info(followers, "followers") // length: 5, capacity: 5

	// add more items beyond slice length but within capacity
	// reflect.Append(s Value, x ...Value) Value
	// Append is a variadic function which means x can be more than one arguments
	// Value is the reflection interface to a Go value
	for i := 3; i < 5; i++ {
		// panic: runtime error: index out of range [3] with length 3
		// followers[i] = fmt.Sprintf("follower #%d", i + 1)

		slice := reflect.ValueOf(followers)
		item := reflect.ValueOf(fmt.Sprintf("beyond length #%d", i+1))
		followers = reflect.Append(slice, item).Interface().([]string)
	}
	print_slice_info(followers, "followers") // length: 5, capacity: 5

	// append multiple items beyond slice length and capacity
	// when you append beyond capacity, golang automatically grows the capacity
	// capacity is for efficiency reasons, best to always create slice with capacity
	// http://xahlee.info/golang/golang_slice.html
	slice := reflect.ValueOf(followers)
	item1 := reflect.ValueOf(fmt.Sprintf("beyond capacity #%d", 101))
	item2 := reflect.ValueOf(fmt.Sprintf("beyond capacity #%d", 102))
	followers = reflect.Append(slice, item1, item2).Interface().([]string)
	print_slice_info(followers, "followers") // length: 7, capacity: 10 (cap is doubled, was 5)

	// A slice, once initialized, is always associated with an underlying array
	// that holds its elements. A slice therefore does not store any data but
	// shares storage with its array and with other slices of the same array.
	//
	// Changing the elements of a slice modifies the corresponding elements of
	// its underlying array. Other slices that share the same underlying array
	// will see those changes.

	// https://golang.org/ref/spec#Slice_expressions
	// slice of slice: slice expression
	// a[low : high]
	//     length = high - low
	//     capacity = high
	// s[a:b] returns a slice of s from index a (included) to b (excluded)
	// s[a:] is same as s[a:len(s)]
	// s[:b] is same as s[0:b]
	var leaders = followers[4:6]
	print_slice_info(leaders, "leaders") // length: 2, capacity: 6

	// full slice expression
	// a[low : high : max]
	//     length = high - low
	//     capacity = max - low
	var leaderCandidates = followers[4:6:8]
	print_slice_info(leaderCandidates, "leader candidates") // length: 2, capacity: 4

	// new slice share the same data with original slice from which it is sliced
	leaders[0] = "vip leader"
	print_slice_info(followers, "followers") // followers[4]: "vip leader"

	// Append more items to slice
	//
	// func append(slice []Type, elems ...Type) []Type
	//     appends elems to the end of slice, and returns the updated slice without modifying the original slice
	//     a new underlying array will be created if there is no sufficient capacity
	//
	// new_slice = append(slice, item1, item2 ...) // ... means any number of items
	// new_slice = append(slice1, slice2...) // slice... is a syntax which means to unpack elements in slice2
	var titles = make([]string, 3, 5)
	var new_titles = append(titles, "principle")
	print_slice_info(titles, "titles")
	print_slice_info(new_titles, "new_titles")

	// append doesn't create a new copy of slice when cap of the
	// result slice is not greater than that of the original slice
	// in this case, new_titles and titles share data from index 0 to 2
	new_titles[1] = "lady gaga"
	print_slice_info(titles, "titles")         // titles[1]: "lady gaga"
	print_slice_info(new_titles, "new_titles") // new_titles[1]: "lady gaga"

	// append multiple items beyond the original slice
	// when cap of the result slice is greater than that of the original slice
	// a whole new slice will be created and it shares no data with the original slice
	// a bigger array is allocated to hold extra elements, and the new slice will point
	// to the newly allocated array
	var all_new_titles = append(titles, "joker", "superman", "batman")
	print_slice_info(titles, "titles")                 // length: 3, capacity: 5
	print_slice_info(all_new_titles, "all_new_titles") // length: 6, capacity: 10

	all_new_titles[2] = "wonder woman"
	print_slice_info(titles, "titles")                 // titles[2]: ""
	print_slice_info(all_new_titles, "all_new_titles") // all_new_titles[2]: "wonder woman"

	// append slice to slice
	fmt.Println("\n#append--------------------------------")

	var s1 = []int{1, 2}
	var s2 = []int{3, 4}
	// note the syntax of the second parameter
	// s2... means to unpack s2 and pass all items of s2 to append
	var s = append(s1, s2...)
	s[0] = 100
	s[2] = 200
	fmt.Printf("%#v\n", s)  // []int{100, 2, 200, 4}
	fmt.Printf("%#v\n", s1) // []int{1, 2}
	fmt.Printf("%#v\n", s2) // []int{3, 4}

	// cut slice (delete elements)
	fmt.Println("\n#cut--------------------------------")

	var sss = []byte("0123456789")
	var ss = append(sss[:5], sss[7:]...) // index 5 is not included, 7 is included
	fmt.Printf("%c\n", ss)               // [0 1 2 3 4 7 8 9]

	// copy slice
	fmt.Println("\n#copy--------------------------------")

	// the number of elements copied is min(len(dst), len(src))
	// the copied elements will in order replace the elements in dst starting at index 0
	var src = []string{"bmw", "audi", "mercedes"}
	var dst = []string{}

	// nothing copied
	copy(dst, src)
	print_slice_info(dst, "dst") // length: 0, capacity: 0 (no elements copied as dst is empty)

	// copy longer to shorter
	dst = make([]string, 2) // make a slice of length 2
	copy(dst, src)
	fmt.Printf("%#v\n", dst) // []string{"bmw", "audi"}

	// copy shorter to longer
	dst = make([]string, 4, 7) // make a slice of length 4 and capacity 7
	copy(dst, src)
	fmt.Printf("%#v\n", dst) // []string{"bmw", "audi", "mercedes", ""}

	// clean slice
	fmt.Println("\n#clean--------------------------------")

	dst = nil // this is the recommanded way
	src = src[0:0]
	fmt.Printf("%#v\n", dst) // []string(nil)
	fmt.Printf("%#v\n", src) // []string{}

	// nested slice
	fmt.Println("\n#nest--------------------------------")

	var matrix = [][]int{{1, 4}, {2, 4}, {3, 6}}
	fmt.Printf("%#v\n", matrix) // [][]int{[]int{1, 4}, []int{2, 4}, []int{3, 6}}

	var haha = make([][]byte, 3) // make a slice with len 3, each element is itself a slice of type [][]byte
	fmt.Printf("%#v\n", haha)    // [][]uint8{[]uint8(nil), []uint8(nil), []uint8(nil)}
	haha[0] = []byte("you")
	haha[1] = []byte("are")
	haha[2] = []byte("beatiful")
	fmt.Printf("%c\n", haha)       // [[y o u] [a r e] [b e a t i f u l]]
	fmt.Printf("%c\n", haha[1][1]) // r

	// slice to string
	// strings.Join(slice, sep string)
	fmt.Println("\n#slice to string--------------------------------")

	var names = []string{"Donald", "Trump"}
	var name = strings.Join(names, " ")
	fmt.Printf("%#v\n", name) // "Donald Trump"
}
