package main

import "fmt"

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

func main() {
	sayHelloTo("Piscki", "Pratama")
	sayHelloTo("John", "Doe")
	sayHelloTo("Fenty", "Rose")
}
