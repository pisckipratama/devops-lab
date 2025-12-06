package main

import "fmt"

func getFullName() (string, string) {
	return "Piscki", "Pratama"
}

func main() {
	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)

	_, namaBelakang := getFullName()
	fmt.Println(namaBelakang)
}
