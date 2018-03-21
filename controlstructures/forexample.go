package main

import "fmt"

func main() {
	s := "囲碁 is the game of Go."
	for index, element := range s {
		fmt.Print(index, ":", string(element), ", ")
	}
}
