package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Server main entry
	staticPage := http.FileServer(http.Dir("./static"))
	http.Handle("/", staticPage) //automatically get homepage
	http.HandleFunc("/about", about)

	fmt.Println("listening on port :8080")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func about(w http.ResponseWriter, r *http.Request) {

}
