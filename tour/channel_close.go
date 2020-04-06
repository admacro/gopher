// https://golang.google.cn/ref/spec#Channel_types
package main

import (
	"fmt"
)

// A sender can close a channel to indicate that no more values will be sent.
// Only the sender should close a channel, never the receiver.
// Channels aren't like files; you don't usually need to close them.
// Closing is only necessary when the receiver must be told there are no more
// values coming, such as to terminate a range loop.
func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)

	// receive values from the channel repeatedly until it is closed.
	// note the syntax here: there is no <- operator, i is same as <-c
	// in this sense, channels act as first-in-first-out queues.

	// iteration over channel permits only one variable, thus it's illegal to write
	//     for i, ok := range c {...}
	for i := range c {
		fmt.Printf("%v ", i) // the last value returned is 0
	}
	fmt.Println()

	// panic: send on closed channel
	// go fibonacci(cap(c), c)

	// Receivers can test whether a channel has been closed by assigning a second
	// parameter to the receive expression.
	//   v, ok := <-ch
	// ok is false if there are no more values to receive and the channel is closed.
	ch := make(chan int, 5)
	go fibonacci(cap(ch), ch)
	for {
		v, ok := <-ch
		if ok {
			fmt.Printf("%v ", v)
		} else {
			fmt.Printf("\nChannel %#v is closed. (Value received from channel: %v)\n", ch, v)
			break
		}
	}
}
