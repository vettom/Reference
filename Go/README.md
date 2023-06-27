# learning Golan
- Is compiled language
-

## simple Data types
- string
	- Interpreted : Can add formatting with in double quotes "I am \n good "
	- Raw string : defind with in back tick `raw text ` What you see is what you get.
- numbers
	- Integer : +ve or -ve numbers. "0, 23, -3"
	- Unsigned integer : minimum value is 0 "6 1 0 145"
	- floating point :  float32 or 64. "4.23e2 4.3"
	- Complex numbers. 
- booleans 
- errors
	- nil value represent no errors

# Variables
Variable declared with type, or go can infer
```go
	var name string //declare a variable
	var name = "vettom" //initialise with value
	var name = "vettom" // with inferred type
	name := "vettom" // Short declaration syntax
	a,b = 2,4 // Multy variable assignment
	var a, b = 6, "Hello" // Different type of values can be declared without type
  	c, d := 7, "World!"

 // Variable can be configured in block
   var (
    a int
    b int    = 1
    c = "hello"
  )
```
# Difference Between var and :=
- var can be used global, however := only with in functions
- Var can declare variable with out value

## print vs println
- println will add whitespaces between arguments and new line at end
## print ()
%v is used to print the value of the arguments
%T is used to print the type of the arguments
```go
	fmt.Printf("i has value: %v and type: %T\n", i, i)
	  fmt.Printf("j has value: %v and type: %T", j, j)
```

### Type conversion ans constant
Go does not allow mixed type and type conversion can be used to change type
Constant is a variable whose value does not change. Constant can be assigned to variable.
```go
	a := 12.352
	var b = int(a) 
	fmt.Println(b)

	const a = 234
	const b string = "string contstant"
	const (
		c = 2
		d = "string value"
	)
```

## Pointers and values
When assigning a value of variable a to  b, value is copied to b. Subsequently if a changes, value of b does not. 
Pointer is when variable point to location of a. So when a changes, value of b also changes. Simply printing value of b prints memory location, not actual value. Use *b to get value
```go
	// value is memory address of a
	x := 20
	y := &x // y is assigned memory address of x
	fmt.Println(y) //print memory address of x
	fmt.Println(*y) //print value of x
	```

