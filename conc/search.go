// https://talks.golang.org/2012/concurrency.slide#42
package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Result string
type Search func(query string) Result

func fakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result(fmt.Sprintf("%v result for %v.", kind, query))
	}
}

var (
	Web   = fakeSearch("Web")
	Image = fakeSearch("Image")
	Video = fakeSearch("Video")
)

func ReplicateSearch(kind string, n int) (replicas []Search) {
	replicas = make([]Search, n)
	for i := range replicas {
		replicas[i] = fakeSearch(fmt.Sprintf("%v-%d", kind, i+1))
	}
	return
}

func Google(query string) (results []Result) {
	c := make(chan Result)
	go func() { c <- First(query, ReplicateSearch("Web", 5)) }()
	go func() { c <- First(query, ReplicateSearch("Image", 3)) }()
	go func() { c <- First(query, ReplicateSearch("Video", 2)) }()

	timeout := time.After(50 * time.Millisecond)
	for i := 0; i < 3; i++ {
		select {
		case result := <-c:
			results = append(results, result)
		case <-timeout:
			fmt.Println("timeout")
			return
		}
	}
	return
}

func First(query string, replicas []Search) Result {
	ch := make(chan Result)
	for i := range replicas {
		replica := replicas[i]
		go func() {
			ch <- replica(query)
		}()
	}
	return <-ch
}

func main() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	results := Google("test")
	elapsed := time.Since(start)
	fmt.Println(results)
	fmt.Println(elapsed)
}
