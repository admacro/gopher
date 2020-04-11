// https://golang.google.cn/ref/spec#Select_statements
package main

import "fmt"

// select chooses one of the cases to run
// the cases are channel communication operations: sends and receives
// only operations that can proceed are valid candidates
// the selection is pseudo-random if more than one cases/operations can proceed
// if no cases/operations can proceed, the default case will be selected
// if no case can proceed and there is no default case, it blocks until one of the cases can proceed
// if in no chance can there be a case that can proceed, it blocks forever

// execution of a select statement
// 1. the right-hand-side of <- of all cases, that is: x -> quit
// 2. see if any of x, quit (or more) channels can proceed (ready to communicate)
// 3. select the case to execute
//    a. if there is only one ready, select that one
//    b. if there are more than one ready, performe a uniform pseudo-random selection
//    c. if there is none ready and there is default case, select the default case
//    d. if there is no default case, the select statement blocks until at least one of the communications can proceed
// 4. execute the communication that's selected
//    a. if the selected case is not a defaut case, the left-hand-side of <- is evaluated
// 5. the statement list of the selected case is executed
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
