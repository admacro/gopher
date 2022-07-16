package main

import (
	"fmt"
	"os"
)

func main() {
	path, err := os.Executable()
	if err != nil {
		panic(err)
	}

	// for scripting languages, script is the code file itself
	// for static languages, it's a compiled binary executable file
	// /var/folders/sp/n3hrxcgd3wgb8zb_3mt1w9kr0000gn/T/go-build931736641/b001/exe/script_path
	fmt.Println(path)
}
