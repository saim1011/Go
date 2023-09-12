package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcom := "Welcome to user input"
	fmt.Println(welcom)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your rating")

	// comma or || err or
	input, _ := reader.ReadString('\n')
	fmt.Println("Thank you ", input)
	fmt.Printf("Data tape is : %T", input)
}
