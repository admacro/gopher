package main

import "fmt"

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

	for i := 0; i < 5; i++ {
		fmt.Println(<-c)
	}

	// receives also block when channel is empty
	// fatal error: all goroutines are asleep - deadlock!
	fmt.Println(<-c)
}
