package main

import "fmt"
import "io/ioutil"
import "os"

func printFileContent(filePath string) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	fmt.Printf("data: %v\n", string(data))
}

func main() {
	var content = []byte("this is a log\nI'm learning golang")
	var filePath = "golang_log.txt"

	// file permission in octal (6 means user can read and write)
	// see more at http://permissions-calculator.org

	// Note
	// if you assign 0644 (octal value of read and write file permission) directly to
	// perm (whose type will be int), ioutil.WriteFile will complain:
	//     cannot use perm (type int) as type os.FileMode in argument to ioutil.WriteFile
	// but if you pass 0644 directly to ioutil.WriteFile, it will work

	// var perm = 0644
	// fmt.Printf("%T\n", perm)				// int

	var perm = os.FileMode(0644)
	err := ioutil.WriteFile(filePath, content, perm)
	// err := ioutil.WriteFile(filePath, content, 0644) // this works as well

	if err != nil {
		panic(err)
	}

	printFileContent(filePath)

	// create file and then write to it step by step
	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}

	var stringContent = "this is a new log\nI'm learning golang again!"
	file.WriteString(stringContent)

	printFileContent(filePath)
}
