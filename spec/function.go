// https://golang.org/ref/spec#Function_types
// https://golang.org/ref/spec#Function_declarations
// https://golang.org/ref/spec#Function_literals
// https://golang.org/doc/effective_go.html#functions
package main

import "fmt"

func processData(fn func(interface{}) string, data ...interface{}) {
	for _, datum := range data {
		fmt.Println(fn(datum))
	}
}

func singleQuote(datum interface{}) (result string) {
	return fmt.Sprintf("'%v'", datum)
}

func doubleQuote(datum interface{}) (result string) {
	return fmt.Sprintf("\"%v\"", datum)
}

// function as type
type processorFunc func(interface{}) string

func quotationFunc(single bool) processorFunc {
	if single {
		return singleQuote
	}
	return doubleQuote
}

func main() {
	processData(singleQuote, 1, 3.14, "你好", "hello", "bonjour")
	processData(doubleQuote, 1, 3.14, "你好", "hello", "bonjour")
}
