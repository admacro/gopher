package main

import "fmt"

// http://xahlee.info/golang/golang_string_backslash_escape.html
func main() {
	// in double quoted string, character sequence starting with
	// backslash may have special meaning. e.g. \n means new line
	//   \t => tab
	//   \r => carriage return
	//
	// aslo, backslash and double quote need to be escaped:
	//   \\ => backslash
	//   \" => double quote
	var s = "AB\t♥🀄"
	fmt.Printf("s = %#v\n", s)

	// \ooo => o is octal digit
	// \xhh => hh is hexdecimal digit
	// \uhhhh => a Unicode character whose code point can be expressed in 4 hex digits (pad 0 in front)
	// \Uhhhhhhhh => a Unicode character whose code point can be expressed in 8 hex digits (pad 0 in front)
	var es = "\101\x42\t\u2665\U0001f004"
	fmt.Printf("ss = %#v\n", es)

	fmt.Printf("s == es is %v", s == es)
}

