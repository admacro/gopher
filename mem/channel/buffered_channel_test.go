package channel

import (
	"fmt"
	"testing"
	"time"
)

func TestBufferedChannel(t *testing.T) {
	bc := NewBc(5)

	for i := 0; i < 5; i++ {
		bc.Send(i)
	}
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		fmt.Println(bc.Receive())
	}
}
