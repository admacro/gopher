package main

import "fmt"
import "path/filepath"
import "os"


func main() {

	// path is root/file.ext
	//   root is the first arg of filepath.Walk
	// Hence, if Walk is called with "abc", and xyz is the file name, then
	// "abc/xyz" is passed to path when walkFn (fn) is called, and info is
	// the os.FileInfo for the named path
	var fn = func(path string, info os.FileInfo, err error) error {
		// first check if there is any error
		if err != nil {
			fmt.Printf("Error %v at path %v\n", err, path)
			return err
		}

		filename := info.Name()
		if info.IsDir() {
			if filename == ".git" {		// ignore git metadata
				return filepath.SkipDir
			}
			filename += "/"
		}

		// print with padding
		fmt.Printf("%v %5v bytes %40v %10v %10v %v\n",
			info.Mode(), info.Size(), info.ModTime(),
			filepath.Ext(path), filepath.Dir(path), filename)

		// return values of walkFn must be one of the following
		// 1. value of type error (error is in interface)
		//    if error is returned, filepath.Walk will also return error and the walk is stopped
		// 2. nil (so far so good, keey walking :-)
		// 3. filepath.SkipDir (value: "skip this directory", type: *errors.errorString)
		//    if filepath.SkipDir is returned, filepath.Walk will skip walking the path
		return nil
	}

	dir := ".."
	fmt.Printf("%v %5v bytes %40v %10v %10v %v\n",
		"Permission", "Size", "Modification Time",
		"Extention", "Directory", "Name")

	// func Walk(root string, walkFn WalkFunc) error
	err := filepath.Walk(dir, fn)
	if err != nil {
		fmt.Printf("Error walking filepath %v\n", err, dir)
	}
}
