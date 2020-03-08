// -----------------------------------------------------------------------------
// [Go Concurrency Patterns](https://talks.golang.org/2012/concurrency.slide#32)
// Author: Rob Pike
// -----------------------------------------------------------------------------

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := fanIn(boring("Jack!"), boring("Joe!"))

	timeout := time.After(1 * time.Second)
	showTime := time.After(5 * time.Second)
	for {
		select {
		case msg := <-c:
			fmt.Printf("You say: %q\n", msg)
		case <-timeout:
			fmt.Println("You are too slow!")
		case <-showTime:
			fmt.Println("You guys talk too much!")
			return
		}
	}
}

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case i := <-input1:
				c <- i
			case i := <-input2:
				c <- i
			}
		}
	}()
	return c
}

func boring(msg string) chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(2e3)) * time.Millisecond)
		}
	}()
	return c
}
