package main

import (
	"fmt"
	"strings"
)

// This function split all strings into all substrings separated by the separator

// syntax:
// func Split(str , sep string) []string

// main func
func main() {
	// initilizing strings

	stry1 := "Welcome, to the, Our Coding, Hollydays"

	stry2 := "My dog name is Derby"

	// display original strings
	fmt.Println("String 1: ", stry1)
	fmt.Println("String 2:", stry2)

	// Display splited strings
	// Using the split func
	splited1 := strings.Split(stry1, ",")
	splited2 := strings.Split(stry2, " ")

	// print the splited strings
	fmt.Println("\nResult of splited string")
	fmt.Println("String 1:", splited1)
	fmt.Println("String 2:", splited2)
}
