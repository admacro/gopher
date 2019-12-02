package main

import "fmt"
import "math/rand"

func main() {
	// math.Seed(seed int64)
	// Seed uses the provided seed value to initialize the default Source
	// to a deterministic state.
	// If Seed is not called, the generator behaves as if seeded by Seed(1).

	var seed int64 = 111						// type of seed is int64
	rand.Seed(seed)
	fmt.Printf("Set seed to %v\n", seed)
	fmt.Printf("rand.Intn(10) => %v\n", rand.Intn(10))		 // 0
	fmt.Printf("rand.Intn(180) => %v\n", rand.Intn(180))	 // 80
	fmt.Printf("rand.Intn(5000) => %v\n", rand.Intn(5000)) // 2420
	for i := 0; i < 30; i++ {
		fmt.Printf("%v", rand.Intn(4))
	} // 201300121330020120220100021311

	// Seed values that have the same remainder when divided by 2^31-1 generate
	// the same pseudo-random sequence.
	// Note: `^` in `2^31 - 1` is not the bitwise operator XOR in a programming
	// context, but the power sign in a mathematical context, thus `2^31 - 1` is
	// not 28, but 2147483647
	// see details here: https://github.com/golang/go/issues/35920
	seed = 2147483647 + 111
	rand.Seed(seed)
	fmt.Printf("\n\nSet seed to %v\n", seed)
	fmt.Printf("rand.Intn(10) => %v\n", rand.Intn(10))		 // 0
	fmt.Printf("rand.Intn(180) => %v\n", rand.Intn(180))	 // 80
	fmt.Printf("rand.Intn(5000) => %v\n", rand.Intn(5000)) // 2420
	for i := 0; i < 30; i++ {
		fmt.Printf("%v", rand.Intn(4))
	} // 201300121330020120220100021311
}
