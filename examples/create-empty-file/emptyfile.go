// Program to create empyt file
package main

import (
	"log"
	"os"
)

func main() {
	// function to create empty file
	// create function can delete file and recreate it if it already eixst or the file will be created
	myfile, err := os.Create("/home/analyst/myprojects/master_Go/examples/create-empty-file/data.txt")

	if err != nil {
		log.Fatal(err)
	}
	log.Println(myfile)
	myfile.Close()
}
