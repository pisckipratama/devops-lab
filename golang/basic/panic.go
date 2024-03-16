package main

import "fmt"

/**
- panic function adalah function yang bisa kita gunakan untuk menghentikan program
- panic function biasanya dipanggil ketika terjadinya panic saat program kita berjalan
- saat panic function dipanggil, program akan terhenti, namun defer function tetap akan dieksekusi
*/

func endApp() {
	fmt.Println("End App")
}

func runApp(error bool) {
	defer endApp()

	if error {
		panic("ERROR!!!")
	}
}

func main() {
	runApp(true)
}
