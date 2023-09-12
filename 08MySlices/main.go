package main

import "fmt"

func main() {
	fmt.Println("Welcome to Slices Class")

	var subList = []string{"Math", "Science", "Tomato", "Peach"}
	fmt.Printf("The type of subList is : %T\n", subList)
	fmt.Println("Subject list: ", subList)
	subList = append(subList, "Physics", "hindi")
	fmt.Println("After Append 2 subject", subList)
	subList = append(subList[1:3])
	fmt.Println("Afer Append colon", subList)
}
