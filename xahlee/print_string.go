package main

import "fmt"

func main() {
	var s = "abc♥好"
	fmt.Printf("s[0] is a byte: %v \n", s[0]) // 97: byte is printed as int by default
	fmt.Printf("s[0] as integer: %d \n", s[0])		// 97: print as integer explicitly
	fmt.Printf("s[0] as hex: %#v \n", s[0])		// 0x61: print as hexdecimal
	fmt.Printf("s[0] as char: %q \n", s[0])		// 'a': print as char

	// print the first byte of ♥ (E2 99 A5)
	fmt.Printf("s[3] as integer: %d \n", s[3]) // 226
	fmt.Printf("s[3] as hex: %#v \n", s[3])		 // 0xe2
	fmt.Printf("s[3] as char: %q \n", s[3])		// 'â'

	// print the last byte of 好 (e5 a5 bd)
	var last_index = len(s) - 1
	fmt.Printf("s[%d] as integer: %d \n", last_index, s[last_index])		// 189
	fmt.Printf("s[%d] as hex: %#v \n", last_index, s[last_index])		// 0xbd
	fmt.Printf("s[%d] as char: %q \n", last_index, s[last_index])		// '½'

	// more printing
	var ss = "abc\t♥\n好"
	fmt.Printf("%s\n", ss)
	//abc	♥
	//好

	// print in golang string syntax, escaping unprintable characters (\t\n)
	fmt.Printf("%q\n", ss)					// "abc\t♥\n好"
	// print in golang string syntax, escaping any unprintable ASCII characters (\t♥\n好)
	fmt.Printf("%+q\n", ss)					// "abc\t\u2665\n\u597d"

	// print in hex, with space between hex numbers
	// if no space between % and x, there will be no space bewteen hex numbers
	// this is best when you want to see how the string is encoded in utf8
	fmt.Printf("% x\n", ss)					// 61 62 63 09 e2 99 a5 0a e5 a5 bd
	fmt.Printf("%x\n", ss)					// 61626309e299a50ae5a5bd

	// print in Unicode notation
	// []rune(ss) converts ss to rune slice (similar to char)
	fmt.Printf("%U\n", []rune(ss)) // [U+0061 U+0062 U+0063 U+0009 U+2665 U+000A U+597D]
}

