package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Read input and print in upper:")
	in := bufio.NewReader(os.Stdin) // Read std in in to buffer
	s, _ := in.ReadString('\n')     //Read string up to the delimeter specified.
	// Second value '_' returned is error
	s = strings.TrimSpace(s) // Trim new line char
	s = strings.ToUpper(s) // Convert string to upper
	fmt.Println(s + "!!!")
}
