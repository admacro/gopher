// https://golang.org/doc/effective_go.html#defer
package main

import (
	"fmt"
	"reflect"
)

type funcTracer interface {
	trace(f string) string
	un(f string)
}

type simpleFunc func()

func (simpleFunc) trace(f string) string {
	fmt.Println("entering", f)
	return f
}

func (simpleFunc) un(f string) {
	fmt.Println("leaving", f)
}

type paramFunc func(in ...interface{})

func (paramFunc) trace(f string) string {
	fmt.Println("entering", f)
	return f
}

func (paramFunc) un(f string) {
	fmt.Println("leaving", f)
}

func valuesOfInterfaces(in ...interface{}) (vals []reflect.Value) {
	vals = make([]reflect.Value, len(in))
	for i, intr := range in {
		vals[i] = reflect.ValueOf(intr)
	}
	return
}

func traceFunc(f funcTracer, in ...interface{}) {
	msg := fmt.Sprintf("%v, params: %v", reflect.TypeOf(f).String(), in)
	defer f.un(f.trace(msg))
	reflect.ValueOf(f).Call(valuesOfInterfaces(in...))
}

func main() {
	var a funcTracer = simpleFunc(func() { fmt.Println("in a") })
	var b funcTracer = paramFunc(
		func(in ...interface{}) {
			traceFunc(a)
			fmt.Println("in b")
		})
	traceFunc(b, 1, 3.14)
}
