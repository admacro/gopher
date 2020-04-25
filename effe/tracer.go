// https://golang.org/doc/effective_go.html#defer
package main

import (
	"fmt"
	"reflect"
	"runtime"
)

type funcTracer interface {
	trace(f string, in []interface{}) string
	un(f string, out []interface{})
}

type voidFunc func()

func (voidFunc) trace(f string, in []interface{}) string {
	fmt.Println("entering", f)
	return f
}

func (voidFunc) un(f string, out []interface{}) {
	fmt.Println("leaving", f)
}

type inFunc func(in ...interface{})

func (inFunc) trace(f string, in []interface{}) string {
	fmt.Printf("entering %v, input: %v\n", f, in)
	return f
}

func (inFunc) un(f string, out []interface{}) {
	fmt.Println("leaving", f)
}

type inOutFunc func(in ...interface{}) []interface{}

func (inOutFunc) trace(f string, in []interface{}) string {
	fmt.Printf("entering %v, input: %v\n", f, in)
	return f
}

func (inOutFunc) un(f string, out []interface{}) {
	fmt.Printf("leaving %v, output: %v\n", f, out)
}

func valuesOfInterfaces(in ...interface{}) (vals []reflect.Value) {
	vals = make([]reflect.Value, len(in))
	for i, intr := range in {
		vals[i] = reflect.ValueOf(intr)
	}
	return
}

func interfacesOfValues(vals []reflect.Value) (in []interface{}) {
	in = make([]interface{}, len(vals))
	for i, val := range vals {
		in[i] = val.Interface()
	}
	return
}

func traceFunc(f funcTracer, in ...interface{}) (out []interface{}) {
	// https://stackoverflow.com/questions/7052693/how-to-get-the-name-of-a-function-in-go
	funcName := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	funcType := reflect.TypeOf(f).String()
	funcInfo := fmt.Sprintf("%v<%v>", funcName, funcType)
	f.trace(funcInfo, in)
	vals := reflect.ValueOf(f).Call(valuesOfInterfaces(in...))
	out = interfacesOfValues(vals)
	defer f.un(funcInfo, out)
	return
}

func someVoidFunc()                { fmt.Println("some arbitrary voidFunc") }
func someInFunc(in ...interface{}) { fmt.Println("some arbitrary inFunc") }
func someInOutFunc(in ...interface{}) []interface{} {
	fmt.Println("some arbitrary inOutFunc")
	return []interface{}{"golang", true}
}
func main() {
	// trace named function
	sv := voidFunc(someVoidFunc)
	si := inFunc(someInFunc)
	sio := inOutFunc(someInOutFunc)
	traceFunc(sv)
	traceFunc(si, 1, 3.14)
	traceFunc(sio, 1, 3.14)

	// anonymous func (func literal)
	var a funcTracer = voidFunc(func() { fmt.Println("in a") })
	var b funcTracer = inFunc(
		func(in ...interface{}) {
			traceFunc(a)
			fmt.Println("in b")
		})
	var c funcTracer = inOutFunc(
		func(in ...interface{}) []interface{} {
			traceFunc(b, in...)
			fmt.Println("in c")
			return []interface{}{"go", "1.14"}
		})
	out := traceFunc(c, 1, 3.14)
	fmt.Println(out)
}
