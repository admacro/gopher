package main

import (
	"fmt"
	"math/rand"
	"sync"
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

// P.S.
// gofmt favors the following formats when formatting the channel operator:
//   when sending value to channel (space at both sides of <-): c <- v
//   when receiving value from channel (space at the value side, not the channel side): v := <-c

var counter = 0
var m sync.Mutex
var wg sync.WaitGroup

func sum(n []int, c chan int, sumId int) {
	defer wg.Done()

	fmt.Printf("[%d] Summing %d numbers: %v\n", sumId, len(n), n)
	sum := 0
	for _, v := range n {
		sum += v
	}

	c <- sum // send value sum to channel c

	// guard critical section
	m.Lock()
	counter = counter + 1
	fmt.Printf("[%d] Number of value sent to channel: %v, value: %v\n", sumId, counter, sum)
	m.Unlock()
}

func main() {
	size := 100
	n := generateRandomNumbers(size)

	c := make(chan int)
	step := 10

	sumId := 1
	for i := range n {
		if i%step == 0 {
			start := i
			end := i + step
			go sum(n[start:end], c, sumId)
			sumId++
			wg.Add(1)
		}
	}

	sum := 0
	for i := 0; i < size/step; i++ {
		sum += <-c
	}
	wg.Wait()

	fmt.Printf("%v numbers sent and received through channel\n", counter)
	fmt.Printf("Sum of %v pseudorandom numbers: %v\n", size, sum)
}

func generateRandomNumbers(size int) []int {
	var n = make([]int, size)
	for i := 0; i < size; i++ {
		n[i] = rand.Intn(10)
	}
	return n
}
