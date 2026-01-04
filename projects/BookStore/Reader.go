package main

import (
	"bufio"
	"os"
	"strings"
)

// func to read user inputs

func Reader() string {

	reader := bufio.NewReader(os.Stdin)

	// reading the input
	text, _ := reader.ReadString('\n')

	return strings.TrimSpace(text)
}
