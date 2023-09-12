package main

import "fmt"

func main() {
	fmt.Println("Welcome to Array Class")

	var subList [4]string
	subList[0] = "Apple"
	subList[1] = "Tomato"
	subList[3] = "Peach"
	fmt.Println("Fruie list is :", subList)
	fmt.Println(len(subList))

	var vegList = [3]string{"first", "Second", "3rd"}
	fmt.Println("List are following: ", vegList)
	fmt.Println(len(vegList))

}
