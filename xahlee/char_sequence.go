// http://xahlee.info/golang/golang_char_sequence.html
package main

import "fmt"
import "unicode/utf8"

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

	pf("type of bytes: %T\n", bs) // []uint8

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
	pf("runes: %q\n", rs)

	// decimal, hexdecimal, and char value of ♥
	pf("runes[6]: dec(%v) hex(%x) char(%c)\n", rs[6], rs[6], rs[6])

	pf("type of runes: %T\n", rs) // []int32

	// rune slice ==> string
	pl("\n# rune slice ==> string ----------")
	pf("runes: %v\n", rs)
	pf("runes: %q\n", rs)

	s = string(rs)
	pf("string: %v\n", s)

	// byte slice ==> rune slice
	pl("\n# byte slice ==> rune slice ----------")
	pf("bytes: %v\n", bs)

	// first convert byte slice to string, then covert string to rune slice
	rs = []rune(string(bs))
	pf("runes: %v\n", rs)
	pf("runes: %q\n", rs)

	// decimal, hexdecimal, and char value of ♥
	pf("runes[6]: dec(%v) hex(%x) char(%c)\n", rs[6], rs[6], rs[6])

	pf("type of runes: %T\n", rs) // []int32

	// rune slice ==> byte slice
	pl("\n# rune slice ==> byte slice ----------")
	pf("runes: %v\n", rs)

	// first convert rune slice to string, then covert string to byte slice
	bs = []byte(string(rs))
	pf("bytes: %v\n", bs)
	pf("bytes: %c\n", bs)

	// decimal, hexdecimal, and char of the first byte of ♥
	pf("bytes[6]: dec(%v) hex(%x) char(%c)\n", bs[6], bs[6], bs[6])

	pf("type of bytes: %T\n", bs) // []int32

	// number of characters (runes)
	pl("\n# number of characters ----------")
	pf("char count of string: %v (using utf8.RuneCount([]byte(s)))\n", utf8.RuneCount([]byte(s)))
	pf("char count of string: %v (using utf8.RuneCountInString)\n", utf8.RuneCountInString(s))
	pf("char count of byte slice: %v\n", utf8.RuneCount(bs))
	pf("char count of rune slice: %v\n", len(rs))

	// substring by character index
	pl("\n# substring by character index ----------")
	rs = []rune(s)
	pf("%v\n", rs[5:7])
	pf("%q\n", rs[5:7])

	// given byte index that starts a char, find the index of the char
	pl("\n# find char index by byte index ----------")
	var i = 9 // index of the first byte of 好 in string s

	pf("index of 好 in s: %v\n", utf8.RuneCountInString(s[0:i]))
	pf("index of 好 in rs: %v\n", utf8.RuneCount(bs[0:i]))

	// given random byte index, find the index that starts a char to
	// which the random byte belongs
	pl("\n# find char index by random byte index ----------")
	var charIndex = func(bs []byte, i int) int {
		for j := i; j >= 0; j-- {
			if utf8.RuneStart(bs[j]) {
				return utf8.RuneCount(bs[0:j])
			}
		}
		return 0
	}
	pf("random byte index in byte slice: %v, char index: %v\n", 7, charIndex(bs, 7))   // 6 ♥
	pf("random byte index in byte slice: %v, char index: %v\n", 9, charIndex(bs, 9))   // 7 好
	pf("random byte index in byte slice: %v, char index: %v\n", 15, charIndex(bs, 15)) // 8 🀄

	var charIndexInString = func(s string, i int) int {
		for j := i; j >= 0; j-- {
			if utf8.RuneStart(s[j]) {
				return utf8.RuneCountInString(s[0:j])
			}
		}
		return 0
	}
	pf("random byte index in string: %v, char index: %v\n", 7, charIndexInString(s, 7))   // 6 ♥
	pf("random byte index in string: %v, char index: %v\n", 9, charIndexInString(s, 9))   // 7 好
	pf("random byte index in string: %v, char index: %v\n", 15, charIndexInString(s, 15)) // 8 🀄

	pl("\n# loop through characters in string ----------")

	// i is index (with respect to byte, not character)
	// c is character (which is rune in golang, its value is an integer)
	for i, c := range s {
		pf("s[%v] = %q (%v) (%U) (%T)\n", i, c, c, c, c)
		// s[6] = '♥' (9829) (U+2665) (int32)
		// s[9] = '好' (22909) (U+597D) (int32)
		// s[12] = '🀄' (126980) (U+1F004) (int32)
	}
}
