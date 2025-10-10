package main

import (
	"bufio"
	"fmt"
	"os"
)

// cli tool to read line of text input from os.stdn
func main() {
	// step
	/*
		- get user input using bufio NewReader(os.Stdin)
		-set variable for taking input
		- store the input in variable call text
		- print the text
	*/

	reader := bufio.NewReader(os.Stdin) // var reader

	fmt.Println("Enter line of text")
	text, _ := reader.ReadString('\n') // var to store user inputs

	fmt.Printf("\nYou Ented: %s ", text)

}
