package main

//This program execute nested -structure

// creating  struct

type Author struct {
	name   string
	editor string
	year   int
}

type Book struct {
	Title    string
	author   Author // nested structure
	location string
}

func main() {
	// getting Book details
	shells := Book{
		Title:    "Ananse In the Land of idiot ",
		location: "Techiman Book Shop",
		author:   Author{"Peter John", "PrintX", 2000},
	}

	// print out shells
	details := &shells

	println("Title:,", (*details).Title)
	println("Store Lacated at : ", (*details).location)

	//Printing nested Struct
	println("Author: ", (*details).author.name)
}
