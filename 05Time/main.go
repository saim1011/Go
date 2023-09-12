package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time ")
	presentTime := time.Now()
	fmt.Println("present Time is :", presentTime)
	fmt.Println(presentTime.Format("01-02-2006 15.04.05 Monday"))

	createdDate := time.Date(2020, time.January, 12, 23, 23, 0, 0, time.Now().Location())
	fmt.Println(createdDate.Format("01-02-2006 Monday"))
}
