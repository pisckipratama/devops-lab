package main

import "fmt"

func main() {
	type NoKtp string

	var ktpPiscki NoKtp = "1111111"

	var contoh = "22222222"
	var contohKtpString = NoKtp(contoh)

	fmt.Println(ktpPiscki)
	fmt.Println(contohKtpString)
}
