package main

import "fmt"

func main() {
	// ahorthand way to declare and initialize an array

	// initialize a shorthand array of strings
	fruits := [4]string{"apple", "banana", "grape", "mango"}

	// print the array
	fmt.Println("ShortHand array output: ")
	for i := 0; i < len(fruits); i++ {
		fmt.Println(fruits[i])
	}
}
