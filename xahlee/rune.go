package main

import "fmt"

// http://xahlee.info/golang/golang_rune.html
func main() {
	// rune means Unicode code point (think of it as char in Java)
	// The purpose of rune is similar to character type in some other languages.
	// rune, the word can be think of as:
	//   1. an integer
	//   2. a golang type (an alias to int32)
	//   3. a unicode codepoint
	//   4. a character

	// rune literal
	// is a syntax to represent one Unicode character
	// note the single quote (just like char in Java)
	// so, value of rune type is just a char
	// for a sequence of rune, use slice
	// slice is a variable length array
	// a slice of rune can be converted to string
	var a = 'a'
	var heart = '♥'
	var zhong = '🀄'
	var newline = '\n'
	var heart_u = '\u2665'				// same as '♥'
	var zhong_u = '\U0001f004'		// same as '🀄'

	// print rune in decimal, hex, and unicode notation
	fmt.Printf("dec: %d, hex: %x, unicode notation: %U\n", heart, heart, heart)

	// actual type of a rune is int32
	fmt.Printf("%T\n", a)					// int32

	fmt.Printf("%c\n", a)					// print char as is
	fmt.Printf("%q\n", heart)			// print in golang syntax (rune syntax)
	fmt.Printf("%U\n", zhong)			// print as Unicode notation (with upper case for a-f)

	fmt.Printf("%b\n", newline)		// base 2 (binary)
	fmt.Printf("%o\n", heart_u)		// base 8 (octal)
	fmt.Printf("%d\n", zhong_u)		// base 10 (decimal)
	fmt.Printf("%x\n", zhong_u)		// base 16 (hexdecimal, with lower case for a-f)
}

