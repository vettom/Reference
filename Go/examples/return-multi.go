package main

import "fmt"

// Returing 2 values and selecter numebr of values
// Function accept 2 input and return that many number of object

type list []string

// Accept l type of list and count type int.
//
//	Return 2 values both are off type list
func deal(l list, count int) (list, list) {
	return l[:count], l[count:]
}

func main() {
	brands := list{"Ford", "VW", "Reno", "Honda", "Cupra"}
	// return 3 brands
	selection, restAll := deal(brands, 2)
	fmt.Println(selection, restAll)

}
