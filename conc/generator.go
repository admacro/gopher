// -----------------------------------------------------------------------------
// [Go Concurrency Patterns](https://talks.golang.org/2012/concurrency.slide#17)
// Author: Rob Pike
// -----------------------------------------------------------------------------

package main

import (
	"fmt"
	"time"
)

// channels are first-class values, just like strings or integers.
// thus, they can be assigned and returned
func main() {
	c := boring("boring!")

	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %q\n", <-c)
	}
	fmt.Println("You are boring. I'm leaving.")
}

func boring(msg string) chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Second)
		}
	}()
	return c
}
