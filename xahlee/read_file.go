package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	const FILE_PATH = "./golang.txt"
	// read whole file and print it
	var content, err = ioutil.ReadFile(FILE_PATH)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%T\n", []byte("hello")) // []uint8 (byte slice)
	fmt.Printf("%T\n", content)         // []uint8 (content is a byte slice)

	fmt.Printf("%v\n", content)  // [84 104 101 32 71 111 32 ... ]
	fmt.Printf("%#v\n", content) // []byte{0x54, 0x68, 0x65, ... }
	fmt.Printf("%q\n", content)  // "The Go programming language ... \n\n ... \n\n ..." (printed in one line)

	fmt.Printf("%v\n", string(content)) // "The Go programming language ..." (printed in natural reading format)

	// read first n bytes
	var getHeadBytes = func(filePath string, n int64) ([]byte, int64) {
		file, err := os.Open(filePath)
		if err != nil {
			panic(err)
		}

		// defer
		// file will be closed when the surrounding function (getHeadBytes) returns,
		// either by return statement or panicking
		defer file.Close()

		// Stat returns the FileInfo structure describing file
		// https://golang.google.cn/pkg/os/#File.Stat
		fileInfo, err := file.Stat()
		if err != nil {
			panic(err)
		}

		// A FileInfo describes a file
		// https://golang.google.cn/pkg/os/#FileInfo
		fmt.Printf("Reading file %v, size: %v bytes, mode: %v\n",
			fileInfo.Name(), fileInfo.Size(), fileInfo.Mode())

		var size = n
		if fileInfo.Size() < n { // Size() of FileInfo is int64
			size = fileInfo.Size()
		}
		var data = make([]byte, size, size)

		// here, count is int ()
		// see file_size_type_issue.go for more
		count, err := file.Read(data)
		if err != nil {
			panic(err)
		}
		return data[:count], int64(count)
	}

	data, count := getHeadBytes(FILE_PATH, 600)
	fmt.Printf("Read %v bytes: %v", count, data)
}
