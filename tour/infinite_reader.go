// http://tour.golang.org/methods/22
package main

import "fmt"

type MyReader struct{}

// emits an infinite stream of the ASCII character 'A'
func (r MyReader) Read(b []byte) (n int, err error) {
	var a = "A"
	for i := range b {
		b[i] = a[0]
		n++
	}
	return n, nil
}

func main() {
	var reader MyReader
	b := make([]byte, 8)
	for i := 0; i < 10; i++ {
		n, err := reader.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
	}
}
