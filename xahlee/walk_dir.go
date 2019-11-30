package main

import "fmt"
import "path/filepath"
import "os"


func main() {

	var fn = func(path string, info os.FileInfo, err error) error {
		// first check if there is any error
		if err != nil {
			fmt.Printf("Error %v at path %v\n", err, path)
			return err
		}

		filename := info.Name()
		if info.IsDir() {
			filename += "/"
		}

		// print with padding

		fmt.Printf("%v %5v bytes %40v %10v %10v %v\n",
			info.Mode(), info.Size(), info.ModTime(),
			filepath.Ext(path), filepath.Dir(path), filename)

		return nil
	}

	dir := "."
	fmt.Printf("%v %5v bytes %40v %10v %10v %v\n",
		"Permission", "Size", "Modification Time",
		"Extention", "Directory", "Name")
	err := filepath.Walk(dir, fn)
	if err != nil {
		fmt.Printf("Error walking filepath %v\n", err, dir)
	}
}
