package main

import (
	"fmt"
	"strings"
)

// The split After func split after a specify separator

func main() {
	// initilize strings
	stry1 := "Welcome, to the, online session, Helloeveryone"
	stry2 := "My cat name is puffi"
	stry3 := "I like to play chess"

	// print the original string
	fmt.Println("String 1: ", stry1)
	fmt.Println("String 2: ", stry2)
	fmt.Println("String 3: ", stry3)

	// print the SplitAfter result
	fmt.Println("Result of SplitAfter ")

	reslt1 := strings.SplitAfter(stry1, ",")
	reslt2 := strings.SplitAfter(stry2, "")
	reslt3 := strings.SplitAfter(stry3, "!")
	rest4 := strings.SplitAfter("", "Helloeveryone, Hello")

	fmt.Println("Result 1: ", reslt1)
	fmt.Println("Result 2: ", reslt2)
	fmt.Println("Result 3: ", reslt3)
	fmt.Println("Result 4: ", rest4)

}
