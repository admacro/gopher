// https://golang.org/ref/spec#Labeled_statements
// A labeled statement may be the target of a goto, break or continue statement.
package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}
	i := 0
Loop:
	if a[i]%2 == 0 {
		fmt.Println(a[i])
	}
	i++
	if i < len(a) {
		goto Loop
	}
}
