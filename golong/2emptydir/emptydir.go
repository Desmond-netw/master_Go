package main

import (
	"log"
	"os"
)

var (
	myfile *os.FileInfo
	es     error
)

func main() {
	// make dir with golang

	// os.Mkdir() method
	if er := os.Mkdir("/home/desmond/Documents/testfile/mkdir1", os.ModePerm); er != nil {
		log.Fatal(er)
	}
}
