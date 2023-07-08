package main

import "fmt"

type vettom []string

func main() {
	members := vettom{"Freya", "Maria"}
	members = append(members, "Lunna")
	// Receiver prints if variable is for type vettom.
	members.print()
}

// Receiver function sets up methos for variables
// Receiver is called with variable name and its type, here type is vettom

func (v vettom) print() {
	for x, y := range v {
		fmt.Println(x, y)
	}
}
