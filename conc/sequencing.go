// -----------------------------------------------------------------------------
// [Go Concurrency Patterns](https://talks.golang.org/2012/concurrency.slide#17)
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
		message1 := <-c
		fmt.Println(message1.content)
		message2 := <-c
		fmt.Println(message2.content)
		message1.readReceipt <- true
		message2.readReceipt <- true
	}
	fmt.Println("You both are boring. I'm leaving.")
}

func boring(msg string) <-chan Message {
	c := make(chan Message)
	readReceipt := make(chan bool)
	go func() {
		for i := 0; ; i++ {
			c <- Message{fmt.Sprintf("%s %d", msg, i), readReceipt}
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
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
