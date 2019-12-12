package main

import (
	"fmt"
	"math/rand"
)

// Channels are a typed conduit through which you can send and receive values
// with the channel operator: <-
// The data flows in the direction of the arrow.

// You create channels using make:
//   make(chan Type)
//   Type is the type of the value you can send and receive with the channel chan

// By default, sends (c <- v) and receives (v <-c) block until the other side is
// ready. This allows goroutines to synchronize without explicit locks or
// condition variables.

func sum(n []int, c chan int) {
	sum := 0
	for _, v := range n {
		sum += v
	}
	c <- sum // send value sum to channel c
}

func main() {
	n := generateRandomNumbers()

	c := make(chan int)
	go sum(n[:len(n)/2], c)
	go sum(n[len(n)/2:], c)

	x, y := <-c, <-c // receive value from c and assign to x and y
	fmt.Println(x, y, x+y)

	// Error when all goroutines have completed their work
	// fatal error: all goroutines are asleep - deadlock!
	// x = <-c
}

func generateRandomNumbers() (n [100]int) {
	for i := range n {
		n[i] = rand.Intn(100)
	}
	return
}
