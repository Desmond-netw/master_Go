package main

import (
	"fmt"
	"strings"
)

// Program to illustrate string trim

func main() {
	// creating, initializing strings
	// Using shorthand methods

	stry1 := "!! Welcome to Everyone !!"
	stry2 := "@@This is example of Golang$$"

	// Display strings
	fmt.Println("Strings before the trimming:")
	fmt.Println("String 1:", stry1)
	fmt.Println("String 1:", stry2)

	// Printing trim strings
	// Using Trim() function
	trim1 := strings.Trim(stry1, "!")
	trim2 := strings.Trim(stry2, "@$")

}
