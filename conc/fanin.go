// -----------------------------------------------------------------------------
// [Go Concurrency Patterns](https://talks.golang.org/2012/concurrency.slide#27)
// Author: Rob Pike
// -----------------------------------------------------------------------------

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// type of the resulting channels from the boring(string) calls
	// will be cast from chan to <-chan when the channels are passed
	// to fanIn(chan string, chan string)
	c := fanIn(boring("Jack!"), boring("Joe!"))

	for i := 0; i < 10; i++ {
		fmt.Printf("You say: %q\n", <-c)
	}
	fmt.Println("You both are boring. I'm leaving.")
}

// Fan-in (https://en.wikipedia.org/wiki/Fan-in)
// input1⟶⬊
//          ⟩⟶ c
// input2⟶⬈

// <-chan means the channel is output only
func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)

	// let whoever is read to talk
	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()

	// when c is returned, it will be cast to <-chan
	return c
}

func boring(msg string) chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}
