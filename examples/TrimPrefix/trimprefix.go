package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Hello, World!"
	output := strings.TrimPrefix(str, "Hello,")
	// without TrimPrefix
	fmt.Println("Without Trim Prefix")
	fmt.Println(str)
	// Trimp prefix
	fmt.Println("Trim the prefix (Hello)")
	fmt.Println(output)

}
