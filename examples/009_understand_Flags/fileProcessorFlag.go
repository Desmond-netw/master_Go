package main

import (
	"flag"
)

func main() {
	// Example to understand flags args
	// A basic file pricessor with flags

	// define flgas
	uppercase := flag.Bool("upper", false, "convert to upppercase")
	reverse := flag.Bool("reverse", false, "reverse the text")
	repeat := flag.Bool("repeat", 1, "Number of times to repeat")

	flag.Parse()
}
