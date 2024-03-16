package main

import "fmt"

func getCompleteName() (firstName, middleName, lastName string, age int8) {
	firstName = "Piscki"
	lastName = "Pratama"

	age = 24
	return firstName, middleName, lastName, age
}

func main() {
	a, b, c, d := getCompleteName()
	fmt.Println(a, b, c, d)
}
