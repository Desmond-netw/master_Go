package main

import (
	"fmt"
	"strconv"
)

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
	fmt.Println("Welcome to the Book Inventory")
	fmt.Println("|-----------------------|")

	// getting books details in variable
	fmt.Printf("Book Title:")
	bookTitle := Reader()

	fmt.Printf("Enter  Quanty:")
	bookQty := Reader()

	fmt.Printf("Enter Unity  Price:")
	bookPrice := Reader()

	// author details
	fmt.Printf("Author Name: ")
	authorName := Reader()

	fmt.Printf("Author Editor: ")
	authorEditor := Reader()

	fmt.Printf("Author year: ")
	authorYear := Reader()

	// print brake line
	fmt.Println()

	// Book data
	qty, _ := strconv.Atoi(bookQty)
	book := Book{
		Title:    bookTitle,
		Quantity: qty,
		Price:    bookPrice,
		BookAuthor: Author{
			Name:   authorName,
			Editor: authorEditor,
			Year:   authorYear,
		},
	}

	// Display inventory
	fmt.Println("\nNew Book Successfully Saved")
	fmt.Println("|--------------------------|")
	fmt.Println("Book Details")

	fmt.Println("Title:", book.Title)
	fmt.Println("Quantity:", book.Quantity)
	fmt.Println("Price:", book.Price)

	fmt.Println("Author Name:", book.BookAuthor.Name)
	fmt.Println("Editor:", book.BookAuthor.Editor)
	fmt.Println("Year:", book.BookAuthor.Year)
}
