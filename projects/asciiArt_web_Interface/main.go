package main

import (
	"flag"
	"fmt"
)

/*
Step:
1- get decode flag and encode flag
2- get server flag to start webserver
- multiline can be decoderd and encodered in web interf
3- check for glags and handle encoding
*/
func main() {
	// define flags to handle server and encoding
	decodeflag := flag.Bool("decode", false, "decode input")
	encodeflag := flag.Bool("encode", false, "encode input")
	//multiline := flag.Bool("multiline", false, "Enable multiline mode")
	//serverflag := flag.Bool("server", false, "Display web interface")

	flag.Parse()

	// validate arguments pass on cli
	if *decodeflag && *encodeflag {
		fmt.Println("Error \nCannot use both -decode and -encode")
		return
	}
	// decodeflag

}
