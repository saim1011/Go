package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome to coding world")
	fmt.Println("Please give your rating")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	fmt.Println("Your rating is : ", input)

	// numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	// if err != nil {
	// 	fmt.Print(err)
	// } else {
	// 	fmt.Print("Added 1 to rating :", numRating+1)
	// }

}
