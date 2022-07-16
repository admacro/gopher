package main

import (
	"fmt"
	"os"
)

func fileExist(path string) bool {
	// according to
	//   https://groups.google.com/forum/#!msg/golang-nuts/Ayx-BMNdMFo/4rL8FFHr8v4J
	// you shouldn't check file existence beforehand, because if you want to do something about
	// the file, you will need to open the file. And the file could disappear between checking
	// and opening. So you simply call os.IsNotExist(err) after you open the file and deal
	// with the err there.

	// Using os.Stat to determine whether a path exist or not
	// this is easy enough for the cases where it is required
	_, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		fmt.Printf("File %v does not exist\n", path)
		return false
	}
	fmt.Printf("File %v exist\n", path)
	return true
}

func main() {
	fileExist("pointer.go")
	fileExist("/Users/james/Downloads/Titanic.mp4")
}
