// -----------------------------------------------------------------------------
// [Go Concurrency Patterns](https://talks.golang.org/2012/concurrency.slide#17)
// Author: Rob Pike
// -----------------------------------------------------------------------------

package main

import (
	"fmt"
	"time"
)

// Launch and ignore
// The go statement runs the function as usual, but doesn't make the caller wait.
// It launches a goroutine.
// The functionality is analogous to the & on the end of a shell command.
//
// goroutine is not a thread. (it's userspace greenthread)
// There might be only one thread in a program with thousands of goroutines.
//
// Instead, goroutines are multiplexed dynamically onto multiple OS threads as
// nedded to keep all the goroutines running; so if one should block, such as
// while waiting for I/O, others continue to run.
//
// If you think of goroutine as a very cheap thread, you won't be far off.
//
// https://golang.org/doc/effective_go.html#goroutines
// A goroutine has a simple model: it is a function executing concurrently with
// other goroutines in the same address space. It is lightweight, costing little
// more than the allocation of stack space. And the stacks start small, so they
// are cheap, and grow by allocating (and freeing) heap storage as required.
func main() {
	c := make(chan string)
	go boring("boring!", c)

	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %q\n", <-c)
	}
	fmt.Println("You are boring. I'm leaving.")
}

// A channel in Go connects two goroutines, so they can communicate (by using <-).
// When sending value to a channel (c <- v), the sender will wait for a receiver to be ready.
// When receiving from a channel (v <-c), the receiver will wait for a value to be sent.
// Both the sender and the receiver must be ready for the communication to take place.
// If one is not ready, the other waits for it to be ready.
// Thus, channels both communicate and synchronize (when the communication happens).
// In other words, channels enable synchrous communiction between goroutines.
func boring(msg string, c chan string) {
	for i := 0; ; i++ {
		c <- fmt.Sprintf("%s %d", msg, i)
		time.Sleep(time.Second)
	}
}

// Buffered channels
// By default, channels have no buffer, that's why they block when the other end is not ready.
// Go channels can be created with a buffer.
// Buffering removes synchronization between two goroutines.
// When buffer is full, synchronization takes place between the sender and the channel
// When buffer is empty, synchronization takes place between the receiver and the channel
// Buffering makes channels more like Erlang's mailboxes.

// Erlang's mailbox is a shared-nothing asynchronous message passing system for
// inter-process communication.
// Reference: https://en.wikipedia.org/wiki/Erlang_(programming_language)
