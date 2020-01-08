// https://golang.google.cn/ref/spec#Select_statements
package main

import "fmt"

// select choose one of the cases to run
// the cases are channel communication operations: sends and receives
// only operations that can proceed are valid candidates
// the selection is pseudo-random if more than one cases/operations can proceed
// if no cases/operations can proceed, the default case will be selected
// if no case can proceed and there is no default case, it blocks until one of the cases can proceed
// if in no chance can there be a case that can proceed, it blocks forever

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit: // blocks on every iteration until `quit <- 0` in anonymous goroutine is run
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c) // blocks on every iteration until `c <- x` in fibonacci() is run
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}
