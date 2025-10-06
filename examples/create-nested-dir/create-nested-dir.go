// Creating Nested dir with os.MkdirAll ()

package main

import (
	"log"
	"os"
)

// using the os.mkdir and os.ModeParm
func main() {
	if es := os.MkdirAll("directory/sub1/sub2/sub3", os.ModePerm); es != nil {
		if os.IsNotExist(es) {
			log.Fatal(es)
		}
	}
}
