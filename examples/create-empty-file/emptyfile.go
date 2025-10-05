// Program to create empyt file
package main

import (
	"log"
	"os"
)

func main() {
	// empty file creation
	// using Creation () function
	myfile, err := os.Create("/home/analyst/myprojects/master_Go/examples/create-empty-file/data.txt")

	if err != nil {
		log.Fatal(err)
	}
	log.Println(myfile)
	myfile.Close()
}
