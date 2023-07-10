package main

//  go is pass by value, meaning when data is passed to function  it makes copy of value.
// If updating value with function, must use pointers to update source
// When * is used in type, it means dealing with TYPE pointer.

// THIS BEHAVIOUR IS FOR STRUCT ONLY. WHEN SLICE, IT DOES NOT MATTER
// That is because slice is reference type
/*
Value Types : Int, string, bool, struct
Ref type: slice, pointers,maps, channels, functions

*/

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	name := "Jenn"
	// when & is used, it gets memory address
	memadd := &name
	// * before variable returns value at memory address
	fmt.Println(memadd, *memadd, name)

	// Create person jenn
	jenn := person{
		firstName: "Jenn",
		lastName:  "Vettom",
	}
	jenn.print()

	// Try update firstname by passing pointer
	jennPointer := &jenn
	jennPointer.updateName("Jennie")
	fmt.Println("\n First update")
	jenn.print()
	// When a receiver is configured with type pointer, can pass pointer value of value itself
	jenn.updateName("Jennifer")
	fmt.Println("\n Second Update")
	jenn.print()
}

// Updating like below cause a new copy to be updated not actual  jenn
// func (p person) updateName(newName string) {
// 	p.firstName = newName
// }

// Updating with pointer value.
//
//	In type *person means dealing with type pointer
//	 *pointer means get value at the memory address.
//
// When receiver type is pointer, go will accept memory address or values itself
func (pointer *person) updateName(newName string) {
	(*pointer).firstName = newName

}

// Receiver Function to print persons details
func (p person) print() {
	fmt.Println(p)
}
