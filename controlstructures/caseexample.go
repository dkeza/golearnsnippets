package main

import "fmt"

func main() {
	x := "test"

	// Tested variable is in switch
	switch x {
	case "other":
		fmt.Println("other")
	case "test":
		fmt.Println("test")
	default:
		fmt.Println("something else")
	}

	// Tested variable is in case
	switch {
	case x == "other":
		fmt.Println("other")
	case x == "test":
		fmt.Println("test")
	default:
		fmt.Println("something else")
	}

}
