// https://golang.org/ref/spec#For_statements
// see ../conc/range_var.go for an example of iteration variable issue
// when they are used in goroutine inside the loop
package main

import "fmt"

func main() {
	// For an array, pointer to array, or slice value aps, if at most one
	// iteration variable is present, the range loop produces iteration values
	// from 0 up to len(a)-1 and does not index into the array or slice itself.
	var aps [5]int          // aps can be array, pointer to array, or slice
	for i, _ := range aps { // no non-blank identifier for receiving value
		// aps is never evaluated as there is no need for the elements of aps
		// len(aps) is constant, thus this loop is equivalent to:
		//     for i := 0; i < 5; i++ {...}
		//  or for i := 0; i <= 4; i++ {...}

		// i ranges from 0 to 4
		fmt.Println(i)
	}

	// For a string value, the "range" clause iterates over the Unicode code
	// points in the string starting at byte index 0. On successive iterations,
	// the index value will be the index of the first byte of successive UTF-8-encoded
	// code points in the string, and the second value, of type rune, will be
	// the value of the corresponding code point. If the iteration encounters
	// an invalid UTF-8 sequence, the second value will be 0xFFFD, the Unicode
	// replacement character, and the next iteration will advance a single byte
	// in the string.
	s := "Go语言"
	for i, r := range s {
		fmt.Printf("%d %v %c\n", i, byte(r), r)
	}

	// The iteration order over maps is not specified and is not guaranteed
	// to be the same from one iteration to the next. If a map entry that has
	// not yet been reached is removed during iteration, the corresponding iteration
	// value will not be produced. If a map entry is created during iteration,
	// that entry may be produced during the iteration or may be skipped. The
	// choice may vary for each entry created and from one iteration to the next.
	m := map[int]int{1: 1, 2: 2, 3: 3, 4: 4, 5: 5}
	for k, v := range m {
		if k == 3 {
			delete(m, 4) // m[4] will not be produced if has not reached

			// for all new entries created during iteration, the choice of whether
			// they are produced or skipped, may vary for each of them, and vary
			// from one iteration to the next.
			// the choices for m[6] and m[7] are independent. One of them may be
			// produced and the other skipped. Or they may both be produced or skipped.
			m[6] = 6 // m[6] may be produced or skipped
			m[7] = 7 // m[7] may be produced or skipped
		}
		fmt.Printf("%v: %v\n", k, v)
	}
}
