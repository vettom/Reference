package main

import "fmt"

// Car brands and color choice. Build all available combination

// Type choice to store the options
type choice []string

// When fuction return value, must have type in definition
// Here return value is of type choice which is a list
func allOPtions() choice {
	brands := []string{"Ferrari", "Honda"}
	colors := []string{"White", "Blue", "Black"}

	// Loop through and save it in seection variable of type choice
	selection := choice{}
	for _, brand := range brands {
		for _, color := range colors {
			selection = append(selection, brand+" "+color)
		}
	}
	return selection
}

func (x choice) printList() {
	// Print any list of type choice
	for i, j := range x {
		fmt.Println(i, j)
	}

}
func main() {
	allCars := allOPtions()
	allCars.printList()
	//  print first car
	fmt.Println(allCars[1:4])

}
