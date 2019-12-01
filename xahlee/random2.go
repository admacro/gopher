package main

import "fmt"
import "math/rand"

func main() {
	// make the same calls as in random.go in different order
	for i := 0; i < 30; i++ {
		fmt.Printf("%v", rand.Intn(4))
	} // 133312100023210231123202333023

	fmt.Println()

	fmt.Printf("rand.Intn(180) => %v\n", rand.Intn(180))	 // 121
	fmt.Printf("rand.Intn(5000) => %v\n", rand.Intn(5000)) // 508
	fmt.Printf("rand.Intn(10) => %v\n", rand.Intn(10))		 // 7
}
