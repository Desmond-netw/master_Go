package main

// nested directory creating
import (
	"log"
	"os"
)

func main() {

	// using os.MkdirAll() function
	if es := os.MkdirAll("parent/sub1/sub2/sub3", os.ModePerm); es != nil {
		if os.IsNotExist(es) {
			log.Println("Folders not created")
		}
	}
}
