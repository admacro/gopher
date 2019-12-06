package main

import (
	"fmt"
	"time"
)

// A Go routine is a lightweight thread managed by Go runtime.
// go f(x, y, z) starts a new goroutine running f(x, y, z).

// The evaluation of f(x, y, z) happens in the current goroutine and the
// execution of f(x, y, z) happens in the new goroutine.

// Goroutines run in the same address space, so access to shared memory must be
// synchronized. The `sync` package provides useful primitives, although you
// don't need them much in Go as there is better choices (channel).

func sayHi(from string, to string) {
	fmt.Printf("%v: Hi %v!\n", from, to)
}

func main() {
	newKid := "Tom"
	fmt.Printf("Mrs. Helen: boys and girls, this is %v. Can you say hi to %v?\n",
		newKid, newKid)
	kids := []string{"Jack", "Emma", "Johnny", "Cindy"}
	for _, kid := range kids {
		go sayHi(kid, newKid)
	}

	// hardcoded waiting time (3s)
	// This waiting is necessary. Because in this case the main() almost always
	// exits before the goroutines ever start, and you will be able to see output
	// from the goroutines.

	// Better (Correct) ways to wait for all goroutines to complete before the
	// main program exits are:
	//   1. sync.WaitGroup
	//   2. channel
	time.Sleep(1000 * time.Millisecond)
}
