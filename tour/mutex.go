package main

import (
	"fmt"
	"sync"
	"time"
)

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	// surround the critical section with Mutex#Lock and Mutex#Unlock
	// the two function calls work like `sync(c.mux){}` in Java
	c.mux.Lock()
	c.v[key]++
	defer c.mux.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	// Lock so only one goroutine at a time can access the map c.v.
	// Also, use defer to let Lock and Unlock go together for better code management.
	c.mux.Lock()
	defer c.mux.Unlock()

	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 10000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}
