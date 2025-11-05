package main

import (
	"flag"
	"fmt"
)

func main() {
	// Example to understand flags args
	// A basic file pricessor with flags

	// define flags
	uppercase := flag.Bool("upper", false, "convert to upppercase")
	reverse := flag.Bool("reverse", false, "reverse the text")
	repeat := flag.Bool("repeat", 1, "Number of times to repeat")

	flag.Parse()

	//get non-flag arguments
	args := flag.Arg()

	if len(args) < 1 {
		fmt.Println("Usage: fileprocessor -upper -reverse -repeat=3 <file>")
		flag.PrintDefaults()
		return
	}
}
