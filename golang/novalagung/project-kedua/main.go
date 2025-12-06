package main

import "fmt"

func main() {
	// manifest typing
	var lastName string
	var firstName string = "Piscki"
	lastName = "Damaya"

	sureName := "Aulia"
	sureName = "Febriansyah"

	fmt.Printf("Halo %s %s %s!\n", firstName, lastName, sureName)
	fmt.Println("Halo", lastName, sureName, firstName+"!")

	var city, country, continent string

	city, country, continent = "Bandung", "Indonesia", "Asia"
	fmt.Println(city, country, continent)

	var first, second, third = "satu", "dua", "tiga"
	fmt.Println(first, second, third)

	name, status, occupation := "Piscki", "Single", "Engineer"
	fmt.Println(name, status, occupation)

	// type inference
	_, isFriday, twoPointTwo, say := 1, false, 2.2, "Hello"
	fmt.Println(isFriday, twoPointTwo, say)

	myName := new(string)
	fmt.Println(myName)
	fmt.Println(*myName)
}
