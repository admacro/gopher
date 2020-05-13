// channel package implements some basic features of Go's built-in channnal
package channel

import (
	"fmt"
	"testing"
	"time"
)

func TestOneWriterOneReaderChannel(t *testing.T) {
	c := New()

	// sender (writer)
	go func(ii int) { c.Send(ii) }(1)

	// receiver (reader)
	// receive in the current goroutine
	// same applies to the following
	fmt.Println(c.Receive())
}

func TestMultiWriterOneReaderChannel(t *testing.T) {
	c := New()

	for i := 0; i < 5; i++ {
		go func(ii int) { c.Send(ii) }(i)
	}

	for i := 0; i < 5; i++ {
		fmt.Println(c.Receive())
	}
}

func TestOneWriterMultiReaderChannel(t *testing.T) {
	c := New()

	for i := 0; i < 5; i++ {
		go func() { fmt.Println(c.Receive()) }()
	}

	for i := 0; i < 5; i++ {
		c.Send(i + 5)
	}

	time.Sleep(time.Millisecond)
}

func TestMultiWriterMultiReaderChannel(t *testing.T) {
	c := New()

	for i := 0; i < 5; i++ {
		go func(ii int) { c.Send(ii) }(i + 10)
	}

	for i := 0; i < 5; i++ {
		go func() { fmt.Println(c.Receive()) }()
	}

	time.Sleep(time.Millisecond)
}
