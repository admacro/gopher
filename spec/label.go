// https://golang.org/ref/spec#Labeled_statements
// A labeled statement may be the target of a goto, break or continue statement.
package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}
	i := 0

EvenOddLoop:
	if a[i]%2 == 0 {
		fmt.Printf("%d is even\n", a[i])
	} else {
		fmt.Printf("%d is odd\n", a[i])
	}

	// labels are not block scoped and do not conflict with identifiers that are not labels
	const EvenOddLoop = 1

	// The scope of a label is the body of the function in which it is declared
	// and excludes the body of any nested function.
	// func() {
	// 	goto EvenOddLoop
	// }()

	i++
	if i < len(a) {
		goto EvenOddLoop
	}
}
