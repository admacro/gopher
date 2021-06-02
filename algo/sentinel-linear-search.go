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

func search(nameList []string, target string) int {
	lastIndex := len(nameList) - 1
	last := nameList[lastIndex]
	nameList[lastIndex] = target
	i := 0
	for ; nameList[i] != target; i++ {
	}
	nameList[lastIndex] = last
	if i < lastIndex || nameList[lastIndex] == target {
		return i
	}
	return -1
}
