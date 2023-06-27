package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Posts struct {
	Id    int
	Title string
}

func main() {
	//  Get from URL
	url := "https://jsonplaceholder.typicode.com/posts"
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error: failed to get URL \n")
		return
	}
	defer response.Body.Close()
	// Use IO util to read conent
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Failed to read response")
		return
	}
	//  Unmarshal the data
	var postsJson []Posts
	err = json.Unmarshal(data, &postsJson)
	if err != nil {
		fmt.Println("Error unmarshaling", err)
		return
	}

	for _, post := range postsJson {
		fmt.Printf("ID: %v\n", post.Id)
		fmt.Printf("Title: %s\n", post.Title)
	}
}
