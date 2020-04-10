// https://golang.org/ref/spec#For_statements
package main

import "fmt"

func main() {
	s := "Go语言"
	// For a string value, the "range" clause iterates over the Unicode code
	// points in the string starting at byte index 0. On successive iterations,
	// the index value will be the index of the first byte of successive UTF-8-encoded
	// code points in the string, and the second value, of type rune, will be
	// the value of the corresponding code point. If the iteration encounters
	// an invalid UTF-8 sequence, the second value will be 0xFFFD, the Unicode
	// replacement character, and the next iteration will advance a single byte
	// in the string.
	for i, r := range s {
		fmt.Printf("%d %v %c\n", i, byte(r), r)
	}
}
