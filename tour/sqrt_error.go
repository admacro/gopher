package main

import (
	"fmt"
	"math"
)

// Go program express error state with error values.
// A nil error denotes success, a non-nil error denotes failure.

// The error type is a built-in interface similiar to fmt.String
// It has a Error() method which returns a string when called.

// As with fmt.Stringer, the fmt package looks for the error interface when
// printing values

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	// here, must convert e to float64 first, otherwise infinite loop
	// as when printing e, fmt.Sprintf will look for e's Error() method which
	// itself calls fmt.Sprintf, thus infinite loop.
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (z float64, err error) {
	if x < 0 {
		err = ErrNegativeSqrt(x)
		return 0, err
	}

	if x == 0 {
		return 0, nil
	}

	const l = 0.000001
	if x > 1 {
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
			return z, nil
		}
		z -= d
	}
}

func main() {
	x := float64(7458)
	s, err := Sqrt(x)
	if err != nil {
		panic(err)
	}
	fmt.Printf("sqrt(%v) = %v\n", x, s)

	x = float64(-7458)
	s, err = Sqrt(x)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("sqrt(%v) = %v\n", x, s)
}
