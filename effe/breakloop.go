// https://golang.org/doc/effective_go.html#switch
package main

import "fmt"

func main() {
OuterLoop:
	for i := 1; i < 5; i++ {
		fmt.Printf("Outerloop %v\n", i)
		for j := 1; j < 5; j++ {
			fmt.Printf("Innerloop %v\n", j)
			if i*j > 3 {
				fmt.Println("break outerloop")
				break OuterLoop
			}
			if i+j > 3 {
				fmt.Println("break innerloop")
				break // break out of the surrounding loop
			}
		}
	}
}
