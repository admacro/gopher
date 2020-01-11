// https://talks.golang.org/2012/concurrency.slide#39
package main

import (
	"fmt"
)

func f(left, right chan int) {
	left <- 1 + <-right
}

func main() {
	const n = 10000
	leftmost := make(chan int) // c1
	right := leftmost
	left := leftmost
	for i := 0; i < n; i++ {
		right = make(chan int)
		go f(left, right) // goroutine f(c1, c2), f(c2, c3), ... , f(c10000, c10001)
		left = right
	}
	go func(c chan int) { c <- 1 }(right) // goroutine func(c10001)
	fmt.Println(<-leftmost)
}
