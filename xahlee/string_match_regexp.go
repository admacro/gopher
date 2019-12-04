package main

import (
	"fmt"
	"regexp"
)

func stringMatchAnyRegexp(s string, patterns []string) bool {
	for _, pattern := range patterns {
		// compile and match step by step
		// reg := regexp.MustCompile(pattern)
		// if reg.Match([]byte(s)) {
		// 	return true
		// }

		// simpler
		matched, err := regexp.MatchString(pattern, s)
		if err != nil {
			panic(err)
		}
		if matched {
			return true
		}
	}
	return false
}

func main() {
	var text = `Alice was beginning to get very tired of sitting by her
sister on the bank, and of having nothing to do: once or twice she had
peeped into the book her sister was reading, but it had no pictures or
conversations in it, «and what is the use of a book,» thought Alice «without
pictures or conversation?».`

	var patterns = []string{"Alice", "Helen"}
	var matched = stringMatchAnyRegexp(text, patterns)
	fmt.Printf("Text match any regexp (%q) => %v\n", patterns, matched)
}
