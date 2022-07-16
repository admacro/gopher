package main

import "fmt"

// function syntax
// func fname(param1 type1, param2 type2, ...) return_type {body}
func add(x int, y int) int {
	return x + y
}

// type spec can be omitted (similar as in variable declaration)
func multiply(x, y int) int {
	return x * y
}

// multiple return values
// named return values are treated as local variables declared
// at the top of the function body;
//
// named result parameters
// the return or result "parameters" of a Go function can be given names
// and used as regular variables, just like the incoming parameters.
// when named, return or result parameters are initialized to the zero
// values for their types when the function begins; if the function executes
// a return statement with no arguments, the current values of the result
// parameters are used as the returned values.
func divide(x, y int) (quotient int, remainder int) {
	quotient = x / y
	remainder = x - quotient*y

	// return quotient, remainder

	// "naked" return: returns the named return values
	// use only in short functions as it can harm readabitlity in
	// longer functions
	return
}

// a function can be declared without a body, such a function
// is implemented outside Go, e.g. an assembly routine
func ImplementedExternally(x, y, z int) string

// function is a value
// it can be
//  1. assgined to a variable
//  2. passed as argument to a function,
//  3. returned by a function
var sub = func(x, y int) int {
	return x - y
}

func main() {
	x, y := 1234, 56
	fmt.Printf("%v + %v = %v\n", x, y, add(x, y))      // 1234 + 56 = 1290
	fmt.Printf("%v * %v = %v\n", x, y, multiply(x, y)) // 1234 * 56 = 69104

	var q, r = divide(x, y)
	fmt.Printf("%v / %v = %v (%v)\n", x, y, q, r) // 1234 / 56 = 22 (2)

	fmt.Printf("%v - %v = %v\n", x, y, sub(x, y)) // 1234 - 56 = 1178

	// apply function to value inline
	var b = func(a int) int { return a + 1 }(x)
	fmt.Println(b) // 1235

	// nested function
	var square = func(x int) int {
		return x * x
	}
	fmt.Printf("%v ^ 2 = %v\n", x, square(x)) // 1234 ^ 2 = 1522756

	// A function type denotes the set of all functions with the same parameter and result types
	// Thus, when two functions types have the same parameter and result types, we say they are
	// of the same function type

	// function as argument
	// note the type of function argument
	// it must be the same as the spec of the function passed over
	var apply = func(x, y int, f func(x, y int) int) int {
		return f(x, y)
	}
	fmt.Printf("apply add on %v and %v: %v\n", x, y, apply(x, y, add))
	fmt.Printf("apply sub on %v and %v: %v\n", x, y, apply(x, y, sub))
	fmt.Printf("apply multiply on %v and %v: %v\n", x, y, apply(x, y, multiply))

	// function type
	// when function is used as a type, you can only specify parameters' types,
	// and parameter's names can be omitted.
	// In real projects, it's better to have names for parameters so that the whoever reads
	// the code can have a better understanding of the meaning for that type.
	type OperationFunc func(int, int) int // same as: type OperationFunc func(x, y int) int

	// func variable
	var tempFunc OperationFunc                                 // zero value is nil
	fmt.Printf("tempFunc: %v, type: %T\n", tempFunc, tempFunc) // tempFunc: <nil>, type: main.OperationFunc

	// return a function
	// same as function argument, the return type for returning a function
	// must be the same as the spec of the function returned
	var operation = func(operator rune) OperationFunc {
		switch operator {
		case '+':
			return add
		case '-':
			return sub
		case '*':
			return multiply
		default:
			return nil
		}
	}

	fmt.Printf("type of result of add: %T\n", add)                       // func(int, int) int
	fmt.Printf("type of result of operation('+'): %T\n", operation('+')) // main.OperationFunc

	fmt.Printf("get operation + and apply on %v and %v: %v\n", x, y, apply(x, y, operation('+')))
	fmt.Printf("get operation - and apply on %v and %v: %v\n", x, y, apply(x, y, operation('-')))
	fmt.Printf("get operation * and apply on %v and %v: %v\n", x, y, apply(x, y, operation('*')))

	// unspecified number of arguments (variadic function)
	// the arguments are received as slice type in the function body
	// the parameter must be the last one in the parameter list
	// the parameter is prefixed with `...`
	var sum = func(base int, addends ...int) int {
		var s = base
		for _, a := range addends {
			s = s + a
		}
		return s
	}

	// variadic function can be invoked with zero or more arguments for that parameter
	fmt.Printf("sum: %v\n", sum(100))               // sum: 100
	fmt.Printf("sum: %v\n", sum(100, 23, 992, 511)) // sum: 1626

	// with a suffix `...`, variadic function can be invoked with slice/array for that parameter
	var n = make([]int, 100, 100)
	for i := range n {
		n[i] = i + 1
	}
	fmt.Printf("sum: %v\n", sum(100, n...)) // sum: 5150
}
