// https://golang.org/doc/effective_go.html#defer
// see ./tracer.go for a more versatile tracer
package main

import "fmt"

func trace(f string) string {
	fmt.Println("entering", f)
	return f
}

func un(f string) {
	fmt.Println("leaving", f)
}

func caller() {
	defer un(trace("caller"))
	fmt.Println("in caller")
	callee()
}

func callee() {
	defer un(trace("callee"))
	fmt.Println("in callee")
}

func main() {
	caller()
}
