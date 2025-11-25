// Create file and writ content to it
package main

import (
	"log"
	"os"
)

// func main
func main() {
	// define string to be wrtien to file
	var myString = "How I started learning kooding at kood/sisu"

	// create a file using os.create function
	file, err := os.Create("/home/analyst/myprojects/master_Go/examples/temp/test.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close() // defer is use for cleanup , like closing file after runing

	// write contents to the file
	_, err = file.WriteString(myString)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal(err)
		}
	}
}
