package main

import "fmt"

func main() {
	name := "Kiki"

	if name == "Piscki" {
		fmt.Println("Orang Ganteng")
	} else if name == "Pratama" {
		fmt.Println("Orang Ganteng Kedua")
	} else {
		fmt.Println("Bukan Orang Ganteng")
	}

	if length := len(name); length > 4 {
		fmt.Println("nama terlalu panjang")
	} else {
		fmt.Println("nama sudah cukup")
	}
}
