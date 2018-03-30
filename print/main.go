package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello")
	fmt.Println("Hello, there")
	fmt.Println("Hello", "there")
	fmt.Print("Hi\n")
	fmt.Printf("Hello, %s", "there", "\n")
	s := fmt.Sprintln("One", "Two", "Three")
	fmt.Print(s)

	fmt.Fprint(os.Stdout, "Hi from Fprint")
}
