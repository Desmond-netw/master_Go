package main

import (
	"fmt"
	"strings"
)

func main() {

	input := "[5 #]"
	// remove brackets using strings openBracket and closeBracket
	brackets := strings.Replace(input, "[", "", -1)
	cleanBracket := strings.Replace(brackets, "]", "", -1)

	// printing
	fmt.Printf("Original text %s\n", input)
	// check if the clean contains ""
	fmt.Printf("Cleaned: %s\n", cleanBracket)
}
