package main

import "fmt"
import "regexp"

func main() {
	var text = `Alice was beginning to get very tired of sitting by her
sister on the bank, and of having nothing to do: once or twice she had
peeped into the book her sister was reading, but it had no pictures or
conversations in it, «and what is the use of a book,» thought Alice «without
pictures or conversation?».`

	var exp = "Alice??(#*)||\\"
	var newStr = "Helen"

	// Compile parses a regular expression and returns, if successful,
	// a Regexp object that can be used to match against text.
	var reg, err = regexp.Compile(exp)
	if err != nil {
		fmt.Printf("regexp: %v", reg) // nil (parse failed)
		panic(fmt.Sprintf("%v, expression => %v", err.Error(), exp))
	}

	// MustCompile is like Compile but panics if the expression cannot be parsed.
	// panic: regexp: Compile(`Alice??(#*)||\`): error parsing regexp: trailing backslash at end of expression: ``
	reg = regexp.MustCompile(exp)
	fmt.Printf("regexp: %v\n", reg) // nil (parse failed)

	exp = "Alice"
	reg = regexp.MustCompile(exp)
	fmt.Printf("regexp: %#v\n", reg) // &regexp.Regexp{expr:"Alice", prog:(*syntax.Prog)(0xc0000741e0), ...

	text = reg.ReplaceAllLiteralString(text, newStr)
	fmt.Println(text)
}
