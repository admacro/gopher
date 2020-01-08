// https://golang.google.cn/ref/spec#Channel_types
package main

import "fmt"

// If the capacity is zero or absent, the channel is unbuffered and communication
// succeeds only when both a sender and receiver are ready. Otherwise, the
// channel is buffered and communication succeeds without blocking if the
// buffer is not full (sends) or not empty (receives).

func main() {
	// Channels can be buffered:
	//   make(chan Type, bufferSize)
	c := make(chan int, 5)

	// send five values to channel c, after which the channel
	// c will be full
	for i := 0; i < 5; i++ {
		c <- i
	}

	// sends block when channel is full
	// fatal error: all goroutines are asleep - deadlock!
	// c <- 123

	// receive five values from channel c, after which the channel
	// c will be empty
	for i := 0; i < 5; i++ {
		fmt.Println(<-c)
	}

	// receives also block when channel is empty
	// fatal error: all goroutines are asleep - deadlock!
	// fmt.Println(<-c)
}
