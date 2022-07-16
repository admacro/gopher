package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// make the same calls as in random.go, but in different order
	// compare output with random.go
	for i := 0; i < 30; i++ {
		fmt.Printf("%v", rand.Intn(4))
	} // 133312100023210231123202333023

	fmt.Println()

	fmt.Printf("rand.Intn(180) => %v\n", rand.Intn(180))   // 121
	fmt.Printf("rand.Intn(5000) => %v\n", rand.Intn(5000)) // 408
	fmt.Printf("rand.Intn(10) => %v\n", rand.Intn(10))     // 7
}
