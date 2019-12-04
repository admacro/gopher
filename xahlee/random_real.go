package main

import "fmt"
import "crypto/rand"

func main() {
	// real random numbers or crypotgraphically random numbers
	// can be obtained by crypto/rand.Read which utilizes the
	// underlining random functions/apis of the system
	var b = make([]byte, 10)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	fmt.Println(b) // different results always
}
