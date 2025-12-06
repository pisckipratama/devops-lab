package main

import "fmt"

/**
- interface adalah tipe data Abstract, dia tidak memiliki implementasi langsung
- sebuah interface berisikan definisi-definisi method
- biasanya interface digunakan sebagai kontrak
*/

type HasName interface {
	GetName() string
}

func SayHello(value HasName) {
	fmt.Println("Hello", value.GetName())
}

type Person struct {
	Name string
}

type Animal struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	person := Person{Name: "Piscki"}
	SayHello(person)

	animal := Animal{Name: "Kuda"}
	SayHello(animal)
}
