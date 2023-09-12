package main

import "fmt"

func main() {

	fmt.Println("Welcome to to Pinter class")

	// var ptr *int
	// fmt.Println("Value of pointer default : ", ptr)

	MyNumber := 23
	var ptr = &MyNumber
	fmt.Println("address of Pointer is :", ptr)
	fmt.Println("Value of Pointer is :", *ptr)

	*ptr = *ptr * 23
	fmt.Print("New values of pointer is ", MyNumber)

}
