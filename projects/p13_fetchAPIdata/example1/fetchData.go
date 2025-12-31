package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// API url
var APIURL = "https://jsonplaceholder.typicode.com/posts"
var client *http.Client

type posts struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

// get func
func GetJson(url string, data interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// check hppt status
	if resp.StatusCode != http.StatusOK {
		fmt.Errorf("HTTP error: %s", resp.Status)
	}

	return json.NewDecoder(resp.Body).Decode(data)
}

// post Func
func GetPost() {
	var post []posts

	err := GetJson(APIURL, &post)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	for i, post := range post {
		fmt.Printf("Post Title %2d. %s\n", i+1, post.Title)
	}

}

// fun main
func main() {
	client = &http.Client{Timeout: 10 * time.Second}
	GetPost()
}
