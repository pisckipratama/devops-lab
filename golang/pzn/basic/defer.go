package main

import "fmt"

/**
- defer function adalah function yang bisa kita jadwalkan untuk dieksekusi setelah sebuah function selesai dieksekusi
- defer function akan selalu dieksekusi walaupun terjadi error di function yang dieksekusi
*/

func logging() {
	fmt.Println("selesai memanggil function")
}

func runApplication() {
	defer logging()
	fmt.Println("Run Application")
}

func main() {
	runApplication()
}
