package main

//This program execute nested -structure

// creating  struct

type Author struct {
	name   string
	editor string
	year   string
}

type Book struct {
	author Author // nested structure
	store  string
}
