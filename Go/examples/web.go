package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", Handler)              // Call handler function
	http.ListenAndServe("localhost:8080", nil) // Listen to serve page
}

// Function to handle request
func Handler(w http.ResponseWriter, r *http.Request) {
	f, _ := os.Open("./menu.txt") // Open file
	io.Copy(w, f)                 // Write contents to http.ResponseWriter
}
