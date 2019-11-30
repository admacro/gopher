package main

import "fmt"
import "io/ioutil"
import "os"

func main() {
	var content = []byte("this is a log\nI'm learning golang")
	var filePath = "golang_log.txt"

	// file permission in octal (6 means user can read and write)
	// see more at http://permissions-calculator.org
	// Note: if you assign 0644 (file permission in octal) directly
	// to perm (actul type is int), ioutil.WriteFile will complain:
	//     cannot use perm (type int) as type os.FileMode in argument to ioutil.WriteFile
	// but if you pass 644 directly to ioutil.WriteFilehough, it works

	// var perm = 0644
	// fmt.Printf("%T\n", perm)				// int

	var perm = os.FileMode(0644)

	err := ioutil.WriteFile(filePath, content, perm)
	// err := ioutil.WriteFile(filePath, content, 0644) // this works
	if err != nil {
		panic(err)
	}

	data, err := ioutil.ReadFile(filePath)
	fmt.Printf("data: %v\n",  string(data))
}
