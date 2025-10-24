package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Hello, World!"
	output := strings.TrimPrefix(str, "Hello,")
	fmt.Println(output)
}
