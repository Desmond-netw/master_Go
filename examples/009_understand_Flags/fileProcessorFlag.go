package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	// Example to understand flags args
	// A basic file pricessor with flags

	// define flags
	uppercase := flag.Bool("upper", false, "convert to upppercase")
	reverse := flag.Bool("reverse", false, "reverse the text")
	// repeat := flag.Bool("repeat", 1, "Number of times to repeat")

	flag.Parse()

	//get non-flag arguments
	args := flag.Args()

	if len(args) < 1 {
		fmt.Println("Usage: fileprocessor -upper -reverse -repeat=3 <file>")
		flag.PrintDefaults()
		return
	}

	fmt.Printf("Processing %d file(s): \n", len(args))

	for _, filename := range args {
		content := fmt.Sprintf("content of %s", filename)

		if *uppercase {
			content = strings.ToUpper(content)
		}
		// check for repeat flag
		if *reverse {
			content = reverseString(content)
		}

		fmt.Printf("%s -> %s\n", filename, content)
	}

}

func reverseString(str string) string {
	// reversing strings
	// converts strings to slice of runes
	runes := []rune(str)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	// convert modify runes back to string
	return string(runes)
}
