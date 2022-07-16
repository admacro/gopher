package main

import "fmt"

func main() {
	// Go has only one loop construct, the for loop

	// classic for loop
	for i := 0; i < 3; i++ {
		fmt.Printf("Loop #%d\n", i+1)
	}

	fmt.Println("---------------------")

	// this acts like "while" in other langs
	i := 1
	for i <= 3 {
		fmt.Printf("Loop #%d\n", i)
		i++
	}

	fmt.Println("---------------------")

	// infinite loop
	x := 1
	for {
		fmt.Printf("Loop #%d\n", x)
		x++

		if x > 3 {
			break
		}
	}

	fmt.Println("---------------------")

	// break and continue acts the same like in other langs
	for i := 1; i < 100; i++ {
		if i%2 == 0 {
			fmt.Printf("Loop #%d\n", i)
			continue
		} else if i%17 == 0 {
			break
		}
	}

	fmt.Println("---------------------")

	// for range loop
	// the first is the index, the second is a copy of the element at the index
	names := []string{"Jack", "Helen", "May"}
	for i, name := range names {
		fmt.Printf("Loop #%d: %v\n", i+1, name)
	}

	// blank identifier _ is for variable not needed
	// thus, it can't be used as value
	for _, name := range names {
		fmt.Printf("Name: %v\n", name)
	}

	// if only index is wanted, the second variable can be omitted
	for i := range names {
		fmt.Printf("Index: %v\n", i)
	}
}
