package main

/*
# Map vs Struct
Map									struct
keys must be of same time    		- Strongly toped
values must of same type	 		- Can be mix
Keys are indexed can iterate 		- Not true
Reference typ						- value type (means it creates copy if not used as pointer)
Ideal for closely related values	-
Dont need to know keys				- Must know fields
*/

import "fmt"

func main() {
	// Multiple ways to define map. Define empty maps
	// Map with key and values as strings
	// var myMap map[string]string
	// myMap := make(map[string]string)

	myMap := map[string]string{
		"golf":   "Blue",
		"passat": "silver",
		"jazz":   "silver",
	}

	myMap["estate"] = "Red"
	delete(myMap, "jazz")
	printMap(myMap)
}

func printMap(m map[string]string) {
	for k, v := range m {
		fmt.Println("Car ", k, "of color ", v)
	}
}
