package main

import "fmt"

func main() {
	fmt.Print("This is my first go app \n")
	// Type conversion
	// a := 12.43
	// var c  = int(a) 
	// fmt.Println(c)

	// Math
	// a,b := 2,10
	// fmt.Println(a + b)
	// fmt.Println(a != b)
	
	// Pointer
	// x := 20
	// y := &x 
	// fmt.Println(y) //print memory address of x
	// fmt.Println(*y) //print value of x
	var slice []int //Initialise empty slice
	slice = []int {1,2,4}
	slice = append(slice,5) // Append value at end
	slice[2] = 3
	fmt.Println(slice)
}	