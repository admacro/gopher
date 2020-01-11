// -----------------------------------------------------------------------------
// [Go Concurrency Patterns](https://talks.golang.org/2012/concurrency.slide#29)
// Author: Rob Pike
// -----------------------------------------------------------------------------

package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Message struct {
	content     string
	readReceipt chan bool
}

func main() {
	c := fanIn(boring("Jack"), boring("Joe"))

	for i := 0; i < 5; i++ {
		message := <-c
		fmt.Println(message.content)

		// Now the receiver needs to do something with the message, so
		// the sender (goroutine) has to wait for some time.
		// After the receiver is done with the message, a receipt is
		// sent back to the sender and the sender can then continue to
		// send messages.
		fmt.Println("Copy!")
		nap()
		message.readReceipt <- true
	}
	fmt.Println("You both are boring. I'm leaving.")
}

// now each speaker must wait for a receipt to
// send the next message
func boring(msg string) <-chan Message {
	c := make(chan Message)
	readReceipt := make(chan bool)
	go func() {
		for i := 0; ; i++ {
			c <- Message{fmt.Sprintf("%s: %d", msg, i), readReceipt}
			nap()
			<-readReceipt // wait for read receipt
		}
	}()
	return c
}

// multi-inputs fan-in
func fanIn(inputs ...<-chan Message) <-chan Message {
	c := make(chan Message)
	for i := range inputs {
		input := inputs[i]
		go func() {
			for {
				c <- <-input
			}
		}()
	}
	return c
}

func nap() {
	time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
}
