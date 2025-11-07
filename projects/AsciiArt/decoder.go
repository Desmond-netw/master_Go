package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

// a func that will decode the ascil text
/*
	Purpose: take args as input and decode it
	Step: Fetch content from input
	- set arguments
	- remove unneeded brackets
	-openBracket [
	-closeBrackets ]
	Step2: validate content
	- spacePosition check
	-conv anything before :spacePostion to int


*/
func main() {
	// parse flags
	flag.Parse()

	// Initilize flags for cli arguments
	arg := flag.Args()
	if len(arg) < 1 {
		fmt.Printf("Usage: go run . '[5 #]' ")
		flag.PrintDefaults()
		os.Exit(1)
	}

	// let set input to first arguments
	input := arg[1]

	// iterate through input to get content and check for openBracket and closeBrackets
	// remove open brackets and closebrackest
	for _, contents := range input {
		content := string(contents)
		cleanedContent := removeBrackets(content)
		fmt.Println(cleanedContent)
		return
	}
	return
}

// func to remove open and close brackets
func removeBrackets(input string) string {
	// Replaceall func to remove all brackets
	return strings.ReplaceAll(strings.ReplaceAll(input, "[", ""), "]", "")
}
