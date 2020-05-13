package channel

import (
	"fmt"
	"testing"
	"time"
)

func TestOneWriterOneReader(t *testing.T) {
	c := make(chan int)

	// sender (writer)
	go func(ii int) { c <- ii }(1)

	// receiver (reader)
	// receive in the current goroutine
	// same applies to the following
	fmt.Println(<-c)
}

func TestMultiWriterOneReader(t *testing.T) {
	c := make(chan int)

	for i := 0; i < 5; i++ {
		go func(ii int) { c <- ii }(i)
	}

	for i := 0; i < 5; i++ {
		fmt.Println(<-c)
	}
}

func TestOneWriterMultiReader(t *testing.T) {
	c := make(chan int)

	for i := 0; i < 5; i++ {
		go func() { fmt.Println(<-c) }()
	}

	for i := 0; i < 5; i++ {
		c <- i + 5
	}

	time.Sleep(time.Millisecond)
}

func TestMultiWriterMultiReader(t *testing.T) {
	c := make(chan int)

	for i := 0; i < 5; i++ {
		go func(ii int) { c <- ii }(i + 10)
	}

	for i := 0; i < 5; i++ {
		go func() { fmt.Println(<-c) }()
	}

	time.Sleep(time.Millisecond)
}
