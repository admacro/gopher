package main

import "fmt"
import "strings"

func main() {
	// len(string) returns the number of bytes in string
	var heart = "♥"
	fmt.Printf("Length of %v is %v.\n", heart, len(heart)) // 3
	fmt.Printf("Length of hello is %v.\n", len("hello"))	 // 5
	fmt.Printf("Length of 你好 is %v.\n", len("你好"))	 // 6

	// substring
	// s[n:m] returns a substring of s from index n to m (excluding m)
	var s = "♥ is heart"
	fmt.Printf("%#v\n", s[0:2])		// "\xe2\x99" (first two bytes of ♥)
	fmt.Println(s[0:3])						// ♥ (s[0] to s[3-1])
	fmt.Println(s[4:len(s)])			// is heart

	// index must be within range, from 0 to length of the string
	// fmt.Println(s[3:20]) // runtime error: slice bounds out of range ...
	// fmt.Println(s[-1:3]) // invalid slice index -1 (index must be non-negative)

	// join string
	// use + to join string
	fmt.Println("a" + "b")				// ab

	// Golang has no string interpolation, use Sprintf

	// String functions are in package "strings"
	var puts = fmt.Println
	puts(strings.Contains(s, "♥")) // true
	puts(strings.ContainsAny(s, "♥∑")) // true

	// Finds whether a string contains a particular Unicode code piont
	// rune means Unicode code point (think of it as char in Java)
	// The purpose of rune is similar to character type in some other languages.
	// rune, the word can be think of as:
	//   1. an integer
	//   2. a Golang type (an alias to int32)
	//   3. a unicode codepoint
	//   4. a character
	// More at http://xahlee.info/golang/golang_rune.html
	puts(strings.ContainsRune(s, '♥')) // true

	puts(strings.HasPrefix(s, "♥")) // true
	puts(strings.HasSuffix(s, "♥")) // false
	puts(strings.ToUpper(s))				 // ♥ IS HEART

	// strings.Trim(string, cutset string)
	// anything appear in cutset are removed
	puts(strings.Trim("1.  I'm learning Golang.  ! ", "! .1"))	// I'm learning Golang

	var ss = "lots of ♥♥♥"
	puts(strings.Count(ss, "♥")) // 3
	puts(strings.Index(ss, "♥")) // 8

	puts(strings.Join([]string{"a", "b", "c"}, "->")) // a->b->c
	puts(strings.Split("123|456|789", "|")) // [123 456 789]
}
