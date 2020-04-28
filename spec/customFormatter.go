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

// receiver is value type Computer
// print functions in fmt package will recognize values
// of both Computer and *Computer
// if receiver is pointer type (*Computer), only values
// of type *Computer will be supported
func (comp Computer) String() string {
	startLine := fmt.Sprintf("Computer {")
	return comp.fieldsString(startLine)
}

// receiver is pointer type *Computer
// only values of type *Computer will be supported by
// print functions in fmt package
func (comp *Computer) Format(f fmt.State, c rune) {
	out := ""
	switch c {
	case 'z':
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
