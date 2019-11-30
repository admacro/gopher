package main

import "fmt"
import "os"
import "os/exec"

func main() {
	dir := "/Users/james/prog"
	err := os.Chdir(dir)					// exec.Command("cd") doesn't work
	if err != nil {
		panic(err)
	}

	// build command
	pwd := exec.Command("pwd")
	output, _ := pwd.Output()			// run it, wait, get output
	fmt.Println(string(output))

	// must use original command, alias doesn't work
	args := []string {"-l", "-a"}
	l := exec.Command("ls", args...)
	lsOutput, _ := l.Output()			// run it, wait, get output
	fmt.Println(string(lsOutput))

	// other functions of Command
	// cmd.Run()											// run it, wait for it to finish
	// cmd.Start()											// run it, don't wait (to get result, use err := cmd.Wait())
}
