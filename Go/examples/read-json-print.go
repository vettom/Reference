package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type User struct {
	Name string `json: "name"`
	Age  int
	Girl bool
}

func main() {
	// Open and read file
	data, err := ioutil.ReadFile("./file.json")
	if err != nil {
		log.Fatal("Failed to open file", err)
	}
	// Check valid json or not Will be handled by unmarshal by default
	jsonCheck := json.Valid(data)
	if jsonCheck {
		fmt.Printf("Json is Valid.\n\n")
	} else {
		log.Printf("Json file Invalid\n")
	}

	// Unmarshal data
	var jsonData User
	err = json.Unmarshal(data, &jsonData)
	if err != nil {
		log.Fatal("Failed to unmarshal", err)
	}

	// Print data
	fmt.Printf("Name: %s\nAge %v\n", jsonData.Name, jsonData.Age)
	// Print yes or no for
	if jsonData.Girl {
		fmt.Printf("Girl?: Yes \n")
	} else {
		fmt.Printf("Girl?: No \n")
	}
}
