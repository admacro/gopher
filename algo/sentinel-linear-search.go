// https://en.wikipedia.org/wiki/Linear_search#With_a_sentinel
package main

import "fmt"

func main() {
	var boys = []string{"Jack", "Bob", "Mike", "Paul", "Tom"}
	printResult(search(boys, "Mike"))
	printResult(search(boys, "Tom"))
	printResult(search(boys, "Scott"))
}

func printResult(index int) {
	if index > 0 {
		fmt.Println("Target found at index ", index)
	} else {
		fmt.Println("Target not found")
	}
}

func search(list []string, target string) int {
	lastIndex := len(list) - 1
	last := list[lastIndex]
	list[lastIndex] = target // this guarantees a normal termation of the for loop below
	i := 0
	for ; list[i] != target; i++ {
	}
	list[lastIndex] = last
	if i < lastIndex || list[lastIndex] == target {
		return i
	}
	return -1
}
