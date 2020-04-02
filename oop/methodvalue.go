// https://golang.org/ref/spec#Method_values
// If the expression x has static type T and M is in the method set of
// type T, x.M is called a method value. The method value x.M is a function
// value that is callable with the same arguments as a method call of x.M.
// The expression x is evaluated and saved during the evaluation of the method
// value; the saved copy is then used as the receiver in any calls, which
// may be executed later.
package main

import "fmt"

type Rt struct{ x int }

func (r Rt) Mv(a int) int {
	v := a + r.x
	fmt.Println(v)
	return v
}

func (r *Rt) Mp(b float64) float64 {
	v := b + float64(r.x)
	fmt.Println(v)
	return v
}

type Ri interface {
	Mi(string) string
}

func (Rt) Mi(s string) string {
	ss := s + "!"
	fmt.Println(ss)
	return ss
}

func main() {
	// 1. type method values
	rt := Rt{123}
	rt.Mv(111)

	// rt.Mv is a method value
	mvMethodValue := rt.Mv
	mvMethodValue(111)

	// automatic referencing
	mpMethodValue := rt.Mp // equivalent to: (&rt).Mp
	mpMethodValue(3.14)

	// automatic dereferencing
	rtp := &rt
	rtpMvMethodValue := rtp.Mv // equivalent to: (*rtp).Mv
	rtpMvMethodValue(222)

	// 2. interface method values
	var ri Ri = &rt
	ri.Mi("bonjour")

	// method value of a method of an interface type
	methodValueI := ri.Mi
	methodValueI("bonjour")
}
