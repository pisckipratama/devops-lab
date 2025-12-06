package main

import "fmt"

func main() {
	var positiveNumber uint8 = 89
	var negativeNumber = -129201933

	fmt.Printf("Bilangan positif: %d\n", positiveNumber)
	fmt.Printf("Bilangan negaatif: %d\n", negativeNumber)

	var decimalNumber = 2.62

	fmt.Printf("Bilangan desimal: %f\n", decimalNumber)
	fmt.Printf("Bilangan desimal: %.3f\n", decimalNumber)

	var exist bool = true
	fmt.Printf("exist? %t \n", exist)

	var message = `Nama saya "John Wick"
Salam kenal.
Mari Belajar "Golang".`

	fmt.Println(message)
}
