package main

import "fmt"

func main() {
	// cara 1
	var name string
	name = "Piscki"
	fmt.Println(name)

	name = "pratama"
	fmt.Println(name)

	// cara 2
	var lastName = "john"
	fmt.Println(lastName)

	lastName = "Erick"
	fmt.Println(lastName)

	// cara 3
	middleName := "Doe"
	fmt.Println(middleName)

	middleName = "Hazard"
	fmt.Println(middleName)

	// cara 4
	var (
		address     = "Bandung"
		nationality = "Indonesia"
	)

	fmt.Println(address)
	fmt.Println(nationality)
}
