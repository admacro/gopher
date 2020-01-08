package main

import (
	"fmt"
	"sync"
)

var s = 0
var t = 0
var m sync.Mutex
var wg sync.WaitGroup

func walk() {
	s = s + 1
	wg.Done()
}

func walkOneAfterAnother() {
	m.Lock()
	t = t + 1
	m.Unlock()
	wg.Done()
}

func main() {
	count := 500
	wg.Add(count)
	for i := 0; i < count; i++ {
		go walk()
	}
	wg.Wait()
	fmt.Println(s)

	wg.Add(count)
	for i := 0; i < count; i++ {
		go walkOneAfterAnother()
	}
	wg.Wait()
	fmt.Println(t)
}
