// https://golang.org/ref/spec#Appending_and_copying_slices
package main

import "fmt"

func main() {
	// append string to byte slice
	var b []byte
	bc := append(b, "Hello"...)

	// the source slice is not modified
	fmt.Printf("%c\n", b)
	fmt.Printf("%c\n", bc)

	// to reuse the source variable, assign the result to source
	bc = append(bc, " Golang"...)
	fmt.Printf("%c\n", bc)
	fmt.Printf("%v\n", string(bc)) // convert to string

	// copy string to byte slice
	var des []byte
	count := copy(des, "Hi")
	fmt.Printf("des capacity: %d\n", cap(des))
	fmt.Printf("%d elements copied: %c\n", count, des) // 0 (des has no capacity)

	des = make([]byte, 10)
	count = copy(des, "Hi你好")
	fmt.Printf("des capacity: %d\n", cap(des))
	fmt.Printf("%d elements copied: %q\n", count, des) // 8 elements copied (8 bytes copied, not character)

}
