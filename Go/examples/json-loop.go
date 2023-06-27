package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name   string `json: "name"`
	Age    int    `json:"age"`
	Email  string `json:"email"`
	Gender string
}

func main() {
	jsonData := []byte(`
		[
			{"name":"John Doe","age":30,"email":"johndoe@example.com"},
			{"name":"Jane Smith","age":25,"email":"janesmith@example.com"}
		]
	`)
	var guests []User
	err := json.Unmarshal(jsonData, &guests)
	if err != nil {
		fmt.Printf("Error decoding json %s\n", err)
	}

	//  Adding a new user
	newuser := User{
		Name:   "Jenn",
		Age:    45,
		Email:  "jenn@vettom.uk",
		Gender: "Femail",
	}
	guests = append(guests, newuser)
	for _, guest := range guests {
		fmt.Printf("Name: %s\n", guest.Name)
		fmt.Printf("Age: %v\n", guest.Age)
		fmt.Printf("Gender: %v\n", guest.Gender)
		fmt.Printf("Email: %s\n ----\n", guest.Email)
	}
}
