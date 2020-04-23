// https://golang.org/ref/spec#Selectors
// if the type of x is a defined pointer type and (*x).f is a valid selector
// expression denoting a field (but not a method), x.f is shorthand for (*x).f
package main

import (
	"fmt"
)

type T1 struct{ x int }

func (T1) M1() { fmt.Println("M1") }

type T2 struct{ *T1 }
type S *T2

func main() {
	t2 := T2{T1: &T1{123}}
	fmt.Println(t2.x)
	fmt.Println(t2.T1.x)
	fmt.Println((*t2.T1).x) // (*t2.T1) means (*(t2.T1)), not (*t2).T1 which is illegal
	t2.M1()
	t2.T1.M1()
	(*t2.T1).M1()

	s := &t2
	fmt.Println(s.x)
	fmt.Println((*s).T1.x)
	fmt.Println((*(*s).T1).x)
	s.M1()
	(*s).T1.M1()
	(*(*s).T1).M1()

}
