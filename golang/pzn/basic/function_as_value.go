package main

import "fmt"

func sayGoodGye(name string) string {
	return "Good bye " + name
}

func main() {
	goodbye := sayGoodGye

	fmt.Println(goodbye("Piscki"))
}
