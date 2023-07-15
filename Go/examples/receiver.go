package main

import "fmt"

// Create a type caled list to store list of objects in slice
type list []string

func main() {
	members := list{"Freya", "Maria"}
	members = append(members, "Lunna")
	// Receiver prints if variable is for type list.
	members.print()
	cars := list{"Golf", "Pssat"}
	cars.print()
}

// Receiver function sets up methos for variables.
// If variable type match, it can call the function.
// Receiver is called with variable name and its type, here type is list

func (v list) print() {
	for x, y := range v {
		fmt.Println(x, y)
	}
}
