package main

import "fmt"

func main() {
	name := ""
	age := 0
	fmt.Println("Enter name and age, and press ENTER")
	i, err := fmt.Scan(&name, &age)
	if err != nil {
		fmt.Println("Error:", err.Error())
	} else {
		fmt.Printf("You entered %d parameters, name %s and age %d \n", i, name, age)
	}
}
