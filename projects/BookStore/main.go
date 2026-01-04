package main

// Building book inventory

// Author Struct
type Author struct {
	Name   string
	Editor string
	Year   string
}

// Book struct
type Book struct {
	Title      string
	BookAuthor Author
	Quantity   int
	Price      string
}

// func main to grab the entry of the books data

func main() {
	println("Welcome to the Book Inventory")

}
