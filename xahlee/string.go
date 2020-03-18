package main

import "fmt"

func main() {
	// interpreted string literal
	// rules for literal string in golang is same as in Java
	var heart = "♥" // ♥(U+2665): E2 99 A5
	var msg = heart + " is a heart"
	fmt.Println(msg)
	fmt.Println("\"use backslash \\ for escaping\"\nnew line after quote")

	// raw string literal
	// grave accent char `
	var anything = `anything ~!@#$%^&*(){}?+|\
      including new line
    but the grave accent char` // not even \` can be included
	fmt.Println(anything)

	// Golang string is a sequence of bytes, not characters.
	// Go string can contain any Unicode character, and also byte
	// sequences that is not valid encoding of any Unicode character.
	// In Go, each character is stored as 1 to 4 bytes by utf8 encoding.
	// You can create a string of any byte by using the hexadecimal escape \x
	fmt.Println("\x41" == "A")
	fmt.Println("\xE2\x99\xA5" == heart) // true
	fmt.Println("\xE2\x99" == heart)     // false

	// backslash escapes interpreted as they are in rune literals with the same restrictions.
	// (except that \' is illegal and \" is legal)
	fmt.Println("\u2665" == "\xE2\x99\xA5") // true
	fmt.Println("\x61" == "\141")           // true

	// s[n] returns the nth byte of of string s
	fmt.Printf("%#v\n", heart[0]) // 0xe2
	fmt.Printf("%#v\n", heart[1]) // 0x99
	fmt.Printf("%#v\n", heart[2]) // 0xa5

	// Same as Python: heart[0] => '\xe2'
	// Different with Ruby: heart[0] => ♥
	fmt.Println(heart[0] == msg[0])
	fmt.Println(heart[1] == msg[1])
	fmt.Println(heart[2] == msg[2])
}
