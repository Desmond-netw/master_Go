// Using IsNotExist to check wether file Exist or Not

package main

import (
	"log"
	"os"
)

// create a point to os.FileInfo
var file *os.FileInfo
var err error

func main() {
	// using Stat() to check file exist and file info
	// the Stat() function can return file infor
	// if the file does not exist it will return error msg.

	//------------------
	file, err := os.Stat("/home/analyst/myprojects/master_Go/examples/temp/data.txt") // specify file path

	if err != nil {
		// ------- checking if file exist
		if os.IsNotExist(err) {
			log.Fatal(err)
		}
	}

	log.Println("File exists ")
	log.Println("File details are : ")
	log.Println("File name : ", file.Name())
	log.Println("File size : ", file.Size())
}
