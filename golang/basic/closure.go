package main

import "fmt"

/**
- closure adalah kemampuan sebuah function berinteraksi dengan data-data disekitarnya dalam scope yang sama
- harap gunakan fitur closure ini dengan bijak saat kita membuat aplikasi, karena bisa membingungkan kode program.
*/

func main() {
	counter := 0

	increment := func() {
		fmt.Println("Increment")
		counter++
	}

	increment()
	increment()
	increment()

	fmt.Println(counter)
}
