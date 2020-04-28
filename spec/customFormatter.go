// implements the Formatter interface
// https://golang.org/pkg/fmt/#Stringer
// https://golang.org/pkg/fmt/#Formatter
// https://github.com/golang/go/blob/go1.14.1/src/fmt/print.go#L53
// https://yourbasic.org/golang/fmt-printf-reference-cheat-sheet/
package main

import (
	"fmt"
	"strings"
)

type Computer struct {
	arch string
	cpu  string
	os   string
}

// https://golang.org/doc/effective_go.html#printing
// There is one important detail to understand about this approach, however:
// don't construct a String/Format method by calling Sprintf in a way that
// will recur into your String/Format method indefinitely. This can happen
// if the Sprintf call attempts to print the receiver directly as a string,
// which in turn will invoke the method again. It's a common and easy mistake
// to make.

// receiver is value type Computer
// print functions in fmt package will recognize values
// of both Computer and *Computer
// if receiver is pointer type (*Computer), only values
// of type *Computer will be supported
func (comp Computer) String() string {
	return comp.fieldsString("Computer {")
}

// receiver is pointer type *Computer
// only values of type *Computer will be supported by
// print functions in fmt package
func (comp *Computer) Format(f fmt.State, c rune) {
	out := ""
	switch c {
	case 'z':
		// %p wants a pointer value, it prints the address the pointer holds
		// if use %z, it will recur forever:
		//     runtime: goroutine stack exceeds 1000000000-byte limit
		//     runtime: sp=0xc0200e0388 stack=[0xc0200e0000, 0xc0400e0000]
		//     fatal error: stack overflow
		startLine := fmt.Sprintf("&Computer(%p) {", comp)
		out = comp.fieldsString(startLine)
	default:
		out = "invalid verb"
	}
	n, err := f.Write([]byte(out))
	if n == 0 && err != nil {
		panic(fmt.Errorf("error formatting %s", out))
	}

}

func (comp *Computer) fieldsString(startLine string) string {
	// Sprintf will only call the String method when it wants a string. To prevent
	// infinite recursion when using Sprintf in String or Format, use a format
	// verb that does not want a string value.
	// String will be called infinitely:
	//     runtime: goroutine stack exceeds 1000000000-byte limit
	//     runtime: sp=0xc0200e0340 stack=[0xc0200e0000, 0xc0400e0000]
	//     fatal error: stack overflow
	// fmt.Println(fmt.Sprint(*comp))

	archLine := fmt.Sprintf("%16s: %v,", "Architecture", comp.arch)
	cpuLine := fmt.Sprintf("%16s: %v,", "CPU", comp.cpu)
	osLine := fmt.Sprintf("%16s: %v", "OS", comp.os)
	endLine := "}"
	return strings.Join([]string{startLine, archLine, cpuLine, osLine, endLine}, "\n")
}

func main() {
	c := Computer{
		arch: "x86-64 (AMD64)",
		cpu:  "AMD Opteron 2212",
		os:   "Ubuntu 18.04.4 LTS (Bionic Beaver)",
	}

	fmt.Println(c)         // will call String() of value type
	fmt.Printf("%z\n", &c) // will call Format() of pointer type

	fmt.Printf("%z\n", c) // value type is not supported
	fmt.Println(&c)       // invalid verb
}
