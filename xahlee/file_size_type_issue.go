package main

import "fmt"
import "os"

// I'm just curious to find out whether this works on a 32bit system
// While I'm on a 64bit Darwin system
// Anyone who can spare a 32bit os, please find out.
func main() {
	// A 3GB file on a 32bit system, more than math.MaxInt32 bytes
	// const FILE_PATH = "./FHD_MOVIE.mkv"
	const FILE_PATH = "/Users/james/prog/allez/xahlee/golang.txt"
	file, err := os.Open(FILE_PATH)
	fileInfo, err := file.Stat()
	if err != nil {
		panic(err)
	}

	defer file.Close()

	// fileInfo.Size is int64
	var size = fileInfo.Size() // Btw, is int64 supported on a 32bit system?
	var data = make([]byte, size, size)

	// count is int
	count, err := file.Read(data) // possible int overflow?
	if err != nil {
		panic(err)
	}
	fmt.Printf("FileInfo.Size() = %v\n", size)
	fmt.Printf("Count of bytes read = %v\n", count)
}
