package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) (z float64) {
	if x == 0 {
		return 0
	}

	const l = 0.000001
	if x > 0 {
		z = x / 2
	} else {
		z = 2 * x
	}

	for {
		fmt.Printf("%v: , ", z)
		d := (z*z - x) / (2 * z)
		fmt.Printf("%v: \n", d)
		if math.Abs(d) < l {
			fmt.Printf("%v: ", z)
			return z
		}
		z -= d
	}
}

func main() {
	x := float64(7458)
	fmt.Printf("sqrt(%v) = %v\n", x, Sqrt(x))

	x = float64(0.00004543)
	fmt.Printf("sqrt(%v) = %v\n", x, Sqrt(x))

	x = float64(0)
	fmt.Printf("sqrt(%v) = %v\n", x, Sqrt(x))
}
