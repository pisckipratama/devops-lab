package main

import "fmt"

func main() {
	name := "Piski"

	switch name {
	case "Piski":
		fmt.Println("Hello Piscki")

	case "Pratama":
		fmt.Println("Hello Tama!")
	default:
		fmt.Println("Hai, boleh kenalan ngga?")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama Terlalu Panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama Lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")

	}
}
