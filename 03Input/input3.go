package main

import (
	"fmt"
)

func main() {
	// Prompt the user for input
	fmt.Print("Enter your name: ")

	// Declare a variable to store user input
	var name string

	// Read user input from the standard input (keyboard)
	_, err := fmt.Scan(&name)

	// Display the user's input
	fmt.Println("Hello,", name, "! Nice to meet you.")
}
