// https://go.dev/doc/tutorial/create-module
package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Set properties of the predefined Logger
	log.SetPrefix("greetigs: ") // log entry prefix
	log.SetFlags(0)             // do not print time, source file, and line number

	// Request greeting messages.
	names := []string{"Gladys", "Emma", "Hare"}
	messages, err := greetings.Hellos(names)

	// If an error was returned, print it to the console and exit the program
	if err != nil {
		log.Fatal(err)
	}

	// Print the messages to the console if no error was returned.
	fmt.Println(messages)
}
