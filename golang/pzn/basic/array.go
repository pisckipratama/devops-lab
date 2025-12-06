package main

import "fmt"

func main() {
	var names [3]string
	names[1] = "Piscki"
	names[2] = "Pratama"
	names[0] = "Doe"

	var values = [3]int{
		90,
		80,
	}

	fmt.Println(values)
	fmt.Println(names[0])

	var city = [...]string{"Jakarta", "Bogor", "Depok", "Tangerang", "Bekasi"}
	fmt.Println(city)
	fmt.Println(len(city))

	// array di golang tidak bisa ditambah
}
