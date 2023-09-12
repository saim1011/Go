package main

import "fmt"

func main() {
	fmt.Println("Enter your first name: ")
	var firstname string
	fmt.Scanln(&firstname)

	fmt.Println("Enter your Second name")
	var SecondNmae string
	fmt.Scanln(&SecondNmae)

	fmt.Println("Your first name is : ", firstname)
	fmt.Println("Your Second Nmae is :", SecondNmae)
}
