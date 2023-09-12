package main

import "fmt"

const LoginToken string = "kdjlakjjflkdjlkfj;" // public

func main() {
	var username string = "Madhurendra"
	fmt.Println(username)
	fmt.Printf("Variable is type of : %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is type of : %T \n", isLoggedIn)

	var smalVal uint8 = 255
	fmt.Println(smalVal)
	fmt.Printf("Variable is type of : %T \n", smalVal)

	var smallFloat float64 = 2.5544567676588686546445634646
	fmt.Println(smallFloat)
	fmt.Printf("Variable is type of : %T \n", smallFloat)

	// Default Values
	var DefaultValue int
	fmt.Println(DefaultValue)
	fmt.Printf("Variable is type of : %T \n", DefaultValue)

	// Implicit type

	var website = "https://hacker.com"
	fmt.Println(website)

	// no var style

	numberOfUser := 300
	fmt.Println("Number users are :", numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("The type of variable is %T", LoginToken)
}
