package main

import "fmt"

func main() {
	var person map[string]string = map[string]string{}

	person["name"] = "Piscki"
	person["address"] = "Bandung"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	book := make(map[string]string)

	book["title"] = "Doa harian seorang Programmer"
	book["author"] = "Figo"
	book["dump"] = "wrong"

	fmt.Println(book)

	delete(book, "dump")

	fmt.Println(book)
}
