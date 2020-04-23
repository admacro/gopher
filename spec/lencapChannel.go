// https://golang.org/ref/spec#Length_and_capacity
// len(channel) returns number of elements queued in channel buffer
// cap(channel) returns channel buffer capacity
package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 10)
	done := make(chan int)

	// make sending faster than receiving so that
	// some elements are queued in channel buffer
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(100 * time.Millisecond)
			c <- i
		}
	}()

	go func() {
		for {
			time.Sleep(500 * time.Millisecond)
			select {
			case <-c:
				fmt.Printf("len(c): %v, cap(c): %v\n", len(c), cap(c))
			case <-done:
				break
			}
		}
	}()

	time.Sleep(5 * time.Second)
	done <- 1
}
