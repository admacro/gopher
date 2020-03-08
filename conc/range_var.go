// https://golang.google.cn/ref/spec#For_statements

// See section `For statements with range clause`
// The iteration variables may be declared by the "range" clause using a form
// of short variable declaration (:=). In this case their types are set to
// the types of the respective iteration values and their scope is the block
// of the "for" statement; they are re-used in each iteration. If the iteration
// variables are declared outside the "for" statement, after execution their
// values will be those of the last iteration.

package main

import (
	"fmt"
	"time"
)

func main() {
	a := []int{1, 2, 3}

	// aa is reused for ever iteration
	// so when goroutine runs, aa might have changed
	// bb is a new instance and is only used for the current iteration
	// so when goroutine runs, bb still holds the same value as assigned
	for _, aa := range a {
		bb := aa
		go func() {
			// loopcloure: loop variable aa captured by func literal
			fmt.Printf("aa: %d \n", aa)
			fmt.Printf("bb: %d \n", bb)
		}()
	}
	time.Sleep(2 * time.Second)
}
