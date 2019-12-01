package main

import "fmt"
import "math/rand"
func main() {
	// Seed values that have the same remainder when divided by 2^31-1 generate
	// the same pseudo-random sequence.
	// different seed, same result
	var seed int64 = 57
	rand.Seed(seed)
	fmt.Printf("Set seed to %v\n", seed)
	fmt.Printf("rand.Intn(10) => %v\n", rand.Intn(10))		 // 4
	fmt.Printf("rand.Intn(180) => %v\n", rand.Intn(180))	 // 143
	fmt.Printf("rand.Intn(5000) => %v\n", rand.Intn(5000)) // 1585
	for i := 0; i < 30; i++ {
		fmt.Printf("%v", rand.Intn(4))
	} // 223301111322332122120302211131
}
