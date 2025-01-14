package main

import (
	"example.com/greetings"
	"fmt"
	"log"
	"rsc.io/quote"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())

	// Request a greeting message.
	message, err := greetings.Hello("Kanav")
	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
			log.Fatal(err)
	}

	// If no error was returned, print the returned message
	// to the console.
	fmt.Println(message)
}