package main

import "fmt"

func main() {
	a, b := 1, 1

	// Simple if
	if a == b {
		fmt.Println("Equal")
	}

	// With else
	a, b = 1, 2
	if a == b {
		fmt.Println("Equal")
	} else {
		fmt.Println("Not equal")
	}

	// With else if
	a, b = 2, 1
	if a == b {
		fmt.Println("Equal")
	} else if a > b {
		fmt.Println("a is greater then b")
	} else {
		fmt.Println("Not equal")
	}

}
