// https://golang.org/doc/effective_go.html#pointers_vs_values
// https://golang.org/pkg/fmt/#Fprintf
package main

import (
	"fmt"
)

type ByteSlice []byte

// implements io.Writer interface
// the idea of using Write on a slice of bytes is
// central to the implementation of bytes.Buffer.
func (p *ByteSlice) Write(data []byte) (n int, err error) {
	slice := *p
	l := len(slice)
	if l+len(data) > cap(slice) {
		newSlice := make([]byte, (l+len(data))*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0 : l+len(data)]
	copy(slice[l:], data)
	*p = slice
	return len(data), nil
}

func main() {
	var b ByteSlice
	fmt.Fprintf(&b, "This hour has %d days\n", 7)
	fmt.Printf("%s", string(b))
}
