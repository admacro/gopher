// https://golang.org/ref/spec#Conversions
package main

import "fmt"

func main() {
	// integer constant to string
	i, j := 71, 0x6F
	fmt.Printf("string(%v) = %q\n", i, string(i))
	fmt.Printf("string(%x) = %q\n", j, string(j))

	// rune to integer, string
	r, rr := 'G', 'o'
	fmt.Printf("int(%q) = %v, string(%q) = %q\n", r, int(r), r, string(r))
	fmt.Printf("int(%q) = %v, string(%q) = %q\n", rr, int(rr), rr, string(rr))

	// []byte/[]rune to string
	goBytes := []byte{'G', 0x6f} // equivalent to: []rune{'G', 0x6f}
	fmt.Printf("string([]byte{'G', 0x6f}) = %q\n", string(goBytes))

	// string (with exactly one character) to integer
	ss, sss := "G", "o"
	fmt.Printf("int(%q[0]) = %v\n", ss, int(ss[0])) // type of ss[0]: uint8 (byte)
	fmt.Printf("int(%q[0]) = %x\n", sss, int(sss[0]))

	// convert nil
	fmt.Printf("type: %T\n", nil)
	bn := []byte(nil)
	fmt.Printf("[]byte(nil) = %q (type: %T)\n", bn, bn)
	ps := (*string)(nil)
	fmt.Printf("(*string)(nil) = %v (type: %T)\n", ps, ps)

	// pointer conversion
	type StringP *string
	s := "test string"
	sp := &s
	ssp := (*string)(sp)
	fmt.Println(*ssp)

	// error: cannot convert sp (type *string) to type string
	// sspf := *string(sp) // same as *(string(sp))

	// channel conversion
	// converting c to a function type Ch: Ch(c)
	// Ch and the type of c must be identical
	// parenthesis are mandatory for the conversion to a channel
	// type that only sends;
	// optional for receive and bidirectional channel types
	type LeftChanInt <-chan int
	var cl LeftChanInt
	cc := (<-chan int)(cl)
	fmt.Println(cc)
	type RightChanInt chan<- int
	var cr RightChanInt
	ccr := chan<- int(cr)
	fmt.Println(ccr)

	// error: cannot convert sp (type *string) to type string
	// sspf := *string(sp) // same as *(string(sp))

	// function conversion
	// converting x to a function type F: F(x)
	// F and the type of x must be identical
	type funcVoid func(int)
	type funcRtrn func(int) string

	var funcVoidVal funcVoid = func(p int) {
		fmt.Printf("function void: %v\n", p)
	}
	// F must be parenthesized when there it has no result list
	// parenthesis are mandatory but gofmt can't add them automatically
	// because the form func(a)(x) is a function signature
	// you must add parenthesis to explicitly make a conversion
	var voidFunc = (func(int))(funcVoidVal)
	voidFunc(123)

	var funcRtrnVal funcRtrn = func(p int) string {
		return fmt.Sprintf("function return: %v", p)
	}
	// parenthesis are optional but gofmt adds them automatically
	// this is also correct: func(int) string(funcRtrnVal)
	var rtrnFunc = (func(int) string)(funcRtrnVal)
	fmt.Println(rtrnFunc(456))

}
