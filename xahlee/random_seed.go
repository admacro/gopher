package main

import "fmt"
import "math/rand"

func main() {
	// math.Seed(seed int64)
	// Seed uses the provided seed value to initialize the default Source
	// to a deterministic state.
	// If Seed is not called, the generator behaves as if seeded by Seed(1).
	// Seed values that have the same remainder when divided by 2^31-1 generate
	// the same pseudo-random sequence.

	var seed int64 = 29						// type of seed is int64
	rand.Seed(seed)
	fmt.Printf("Set seed to %v\n", seed)
	fmt.Printf("rand.Intn(10) => %v\n", rand.Intn(10))		 // 3
	fmt.Printf("rand.Intn(180) => %v\n", rand.Intn(180))	 // 36
	fmt.Printf("rand.Intn(5000) => %v\n", rand.Intn(5000)) // 4450
	for i := 0; i < 30; i++ {
		fmt.Printf("%v", rand.Intn(4))
	} // 102003022300333212232000211300
}
