package main

import "fmt"

func main() {
	// konversi angka
	var angka32 = 32769
	var angka64 = int64(angka32)
	var angka16 = int16(angka32)

	fmt.Println(angka32)
	fmt.Println(angka64)
	fmt.Println(angka16)

	// konversi string
	var name string = "Testing"

	var t uint8 = name[0]
	var tString = string(t)

	fmt.Println(name)
	fmt.Println(t)
	fmt.Println(tString)
}
