package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello world")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter you name: ")
	fmt.Println("Your name is :", reader)

}
