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
// at the top of the function body
func divide(x, y int) (quotient int, remainder int) {
	quotient = x / y
	remainder = x - quotient*y

	// return quotient, remainder

	// "naked" return: returns the named return values
	// use only in short functions as it can harm readabitlity in
	// longer functions
	return
}

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

	// function as argument
	// note the type of function argument
	// it must be the same as the spec of the function passed over
	var apply = func(x, y int, f func(x, y int) int) int {
		return f(x, y)
	}
	fmt.Printf("apply add on %v and %v: %v\n", x, y, apply(x, y, add))
	fmt.Printf("apply sub on %v and %v: %v\n", x, y, apply(x, y, sub))
	fmt.Printf("apply multiply on %v and %v: %v\n", x, y, apply(x, y, multiply))

	// return a function
	// same as function argument, the return type for returning a function
	// must be the same as the spec of the function returned
	var operation = func(operator rune) func(x, y int) int {
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
	fmt.Printf("get operation + and apply on %v and %v: %v\n", x, y, apply(x, y, operation('+')))
	fmt.Printf("get operation - and apply on %v and %v: %v\n", x, y, apply(x, y, operation('-')))
	fmt.Printf("get operation * and apply on %v and %v: %v\n", x, y, apply(x, y, operation('*')))

	// unspecified number of arguments (variadic function)
	// the arguments are received as slice type in the function body
	var sum = func(addends ...int) int {
		var s = 0
		for _, a := range addends {
			s = s + a
		}
		return s
	}
	var n = make([]int, 100, 100)
	for i, _ := range n {
		n[i] = i + 1
	}
	// pass slice/array to function
	fmt.Printf("sum 1 to 100: %v\n", sum(n...)) // sum 1 to 100: 5050
}
