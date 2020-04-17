// implements the Formatter interface
// https://golang.org/pkg/fmt/
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

func (comp *Computer) Format(f fmt.State, c rune) {
	startLine := fmt.Sprintf("Computer&%v {", &comp)
	archLine := fmt.Sprintf("%16s: %v,", "Architecture", comp.arch)
	cpuLine := fmt.Sprintf("%16s: %v,", "CPU", comp.cpu)
	osLine := fmt.Sprintf("%16s: %v", "OS", comp.os)
	endLine := "}"
	out := strings.Join([]string{startLine, archLine, cpuLine, osLine, endLine}, "\n")
	n, err := f.Write([]byte(out))
	if n == 0 && err != nil {
		panic(fmt.Errorf("error formatting %s", out))
	}
}

func main() {
	c := Computer{
		arch: "x86-64 (AMD64)",
		cpu:  "AMD Opteron 2212",
		os:   "Ubuntu 18.04.4 LTS (Bionic Beaver)",
	}
	var fc fmt.Formatter = &c
	fmt.Printf("%v\n", fc)
}
