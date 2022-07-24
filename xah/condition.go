package main

import (
	"fmt"
	"runtime"
)

func grade(score int) string {
	if score < 0 || score > 100 {
		// galang's panic is like throwing runtime exception in Java
		panic("Score must be between 0 and 100.")
	} else if score < 60 {
		return "D"
	} else if score < 80 {
		return "C"
	} else if score < 90 {
		return "B"
	} else {
		return "A"
	}
}

// Go only runs the selected case
// break is provided automatically in Go
// switch cases need not be constants, and the values involved
// need not be integers
func gender(code int) string {
	switch code {
	case 0:
		return "Male"
	case 1:
		return "Female"
	case 2:
		return "Transgender"
	default:
		return "Unknown"
	}
}

func main() {
	fmt.Printf("My score is %v.\n", grade(0))
	// fmt.Printf("My score is %v.\n", grade(101)) // error
	fmt.Printf("My score is %v.\n", grade(87))
	fmt.Printf("My score is %v.\n", grade(99))
	fmt.Printf("My score is %v.\n", grade(70))

	n := 4
	if x := 3; n > x { // x := 3 is a short statement inside if clause
		fmt.Println("OK")
	}

	// NO ternary expression
	// beautiful_and_rich := true
	// fmt.Printf("Will you marry me? %v\n", beautiful_and_rich ? "Yes." : "No" )

	// switch case
	fmt.Printf("I'm %v.\n", gender(1))
	fmt.Printf("I'm %v.\n", gender(2))
	fmt.Printf("I'm %v.\n", gender(0))
	fmt.Printf("I'm %v.\n", gender(-1))

	c := 1
	switch c {
	case 0, 1: // multiple tests
		fmt.Println("Binary")
		fallthrough // transfer control to the first statement of the next case
	case 101:
		fmt.Println("Mandatory") // this line will be excuted
	default:
		fmt.Println("Done") // this line will be ignored
	}

	// without switch expression, each case expression just tests for true
	// just like multiple ifelse
	switch { // same as `switch true`
	case c < 1:
		fmt.Println("Less")
	case c == 1:
		fmt.Println("Equal")
	case c > 1:
		fmt.Println("Greater")
	}

	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	// Switch cases evaluate cases from top to bottom, stopping when a case succeeds.
	var f = func() bool {
		fmt.Println("f()")
		return false
	}
	var i int = 0
	switch {
	case i < 2:
	case f(): // does not call f() if i < 2
	}

	// f() will be called as i > 2 is false
	switch {
	case i > 2:
	case f():
	default:
		fmt.Println("default case")
	}
}
