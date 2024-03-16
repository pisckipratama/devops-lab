package main

import "fmt"

func getHello(name string) string {
	hello := "Hello " + name

	return hello
}

func main() {
	fmt.Println(getHello("Piscki"))
}
