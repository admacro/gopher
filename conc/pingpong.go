// -----------------------------------------------------------------------------
// [Advanced Go Concurrency Patterns](https://talks.golang.org/2013/advconc.slide#6)
// Author: Sameer Ajmani
// -----------------------------------------------------------------------------

package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Ball struct{ hits int }

func Player(name string, table chan *Ball) {
	for {
		ball := <-table
		ball.hits++
		fmt.Println(name, ball.hits)
		time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
		table <- ball
	}
}

func main() {
	table := make(chan *Ball)

	go Player("ping", table)
	go Player("pong", table)

	// new(Ball) is same as &Ball{}
	// see https://golang.google.cn/ref/spec#Allocation
	table <- new(Ball) // toss the ball; game on (deadlock if commented out)

	time.Sleep(time.Second)
	<-table // grab the ball; game over
	fmt.Println("Game over")
}
