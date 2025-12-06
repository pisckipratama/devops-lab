package main

import "fmt"

/**
- recover adalah function yang bisa kita gunakan untuk menangkap data panic
- dengan recover proses panic akan terhenti, sehingga program akan tetap berjalan
*/

func finishApp() {
	fmt.Println("End App")

	message := recover()
	fmt.Println("terjadi panic", message)
}

func runApps(error bool) {
	defer finishApp()

	if error {
		panic("Ups Error")
	}
}

func main() {
	runApps(true)
	fmt.Println("Piscki Pratama")
}
