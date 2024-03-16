package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeCountryToIndoneia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	address := Address{}

	ChangeCountryToIndoneia(&address)

	fmt.Println(address)
}
