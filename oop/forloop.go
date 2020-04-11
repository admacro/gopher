// https://golang.org/ref/spec#For_statements
package main

import "fmt"

func main() {
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
