# Arrays Maps

## Array
Arrays have fixed size.
```go
	var arr [3]int //Array with 3 int elements all are 0 value
	arr := [3]int(1,2,3) //Array elements with value
	fmt.Println(arr[0]) // Print first element
	arr[2] = 10 // Update value at element 3
	len(arr) // Length array
```

## Slice 
Is a reference type array that can grow or shrink. It does not hold value of its own, but a pointer to elements in another array. Changes to elements in Array or slice will reflect at both places.
> Slice is when initialised without number of elements []
```go
	slice := []int //Initialise empty slice
	ar := [2]int // Is array with 2 elements

	slice = []int {1,2,4} // Slice with values
	slice = append(slice,5)


```