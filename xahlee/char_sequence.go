// http://xahlee.info/golang/golang_char_sequence.html
package main

import "fmt"

// char sequence in golang can be represented in 3 different formats
//   1. string (immutable byte sequence)
//   2. byte slice (mutable byte sequence)
//   3. rune slice (re-grouping of byte slice so that each element is a char)

// string is a nice way to deal with short sequence (of bytes or characters)
// everytime you "modify" a string, a new string is created as it's immutable
// this is inefficient when the string is huge, such as file content

// byte slice is just like string, but mutable
// you can modify each byte or character
// this is very efficient with file content (txt or binary) or IO stream from networking

// rune slice is like byte slice
// except that each index is a character instead of a byte
// this is best if you work with text files that have lots of non-ASCII characters,
// such as chinese, emoji, or math symbols

var pl = fmt.Println
var pf = fmt.Printf

func main() {
	var s = "abc123♥好🀄"

	// string ==> byte slice
	pl("\n# string ==> byte slice ----------")
	pf("string: %v\n", s)

	var bs = []byte(s)
	pf("bytes: %v\n", bs)
	pf("bytes: %c\n", bs)

	// decimal, hexdecimal, and char value of first byte of ♥
	pf("bytes[6]: dec(%v) hex(%x) char(%c)\n", bs[6], bs[6], bs[6])

	pf("type of bytes: %T\n", bs)	// []uint8


	// byte slice ==> string
	pl("\n# byte slice ==> string ----------")
	pf("bytes: %v\n", bs)
	pf("bytes: %c\n", bs)

	s = string(bs)
	pf("string: %v\n", s)


	// string ==> rune slice
	pl("\n# string ==> rune slice ----------")
	pf("string: %v\n", s)

	var rs = []rune(s)
	pf("runes: %v\n", rs)
	pf("runes: %c\n", rs)

	// decimal, hexdecimal, and char value of ♥
	pf("runes[6]: dec(%v) hex(%x) char(%c)\n", rs[6], rs[6], rs[6])

	pf("type of runes: %T\n", rs)	// []int32


	// rune slice ==> string
	pl("\n# rune slice ==> string ----------")
	pf("runes: %v\n", rs)
	pf("runes: %c\n", rs)

	s = string(rs)
	pf("string: %v\n", s)


	// byte slice ==> rune slice
	pl("\n# byte slice ==> rune slice ----------")
	pf("bytes: %v\n", bs)

	// first convert byte slice to string, then covert string to rune slice
	rs = []rune(string(bs))
	pf("runes: %v\n", rs)
	pf("runes: %c\n", rs)

	// decimal, hexdecimal, and char value of ♥
	pf("runes[6]: dec(%v) hex(%x) char(%c)\n", rs[6], rs[6], rs[6])

	pf("type of runes: %T\n", rs)	// []int32


	// rune slice ==> byte slice
	pl("\n# rune slice ==> byte slice ----------")
	pf("runes: %v\n", rs)

	// first convert rune slice to string, then covert string to byte slice
	bs = []byte(string(rs))
	pf("bytes: %v\n", bs)
	pf("bytes: %c\n", bs)

	// decimal, hexdecimal, and char of the first byte of ♥
	pf("bytes[6]: dec(%v) hex(%x) char(%c)\n", bs[6], bs[6], bs[6])

	pf("type of bytes: %T\n", bs)	// []int32
}
