package channel

import (
	"fmt"
	"testing"
	"time"
)

func TestBufferedChan(t *testing.T) {
	bc := make(chan int, 5)

	for i := 0; i < 5; i++ {
		bc <- i
	}
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		fmt.Println(<-bc)
	}
}
