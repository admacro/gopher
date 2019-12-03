package main

import "fmt"

func main() {
	var generator = func() (f func(i int) (n int)) {
		var a, b = 0, 1
		f = func(i int) (n int) {
			switch i {
			case 0:
				return a
			case 1:
				return b
			default:
				c := a + b
				b, a = c, b
				return c
			}
		}
		return f
	}
	var fib = generator()
	for i := 0; i < 13; i++ {
		fmt.Println(fib(i))
	}
}
