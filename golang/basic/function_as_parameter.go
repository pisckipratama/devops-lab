package main

import "fmt"

type Filter func(string) string //using declaration

func sayHelloWithFilter(name string, filter Filter) {
	filteredName := filter(name)
	fmt.Println("Hello", filteredName)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("Piscki", spamFilter)

	filter := spamFilter
	sayHelloWithFilter("Anjing", filter)
}
