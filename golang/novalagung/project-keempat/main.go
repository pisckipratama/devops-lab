package main

import "fmt"

func main() {
	const (
		square         = "kotak"
		isToday  bool  = true
		numeric  uint8 = 8
		floatNum       = 2.2
	)

	fmt.Println(square, isToday, numeric, floatNum)

	const (
		a = "konstanta"
		b
	)

	fmt.Println(a, b)

	const (
		today string = "monday"
		sekarang
		isToday2 = true
	)

	fmt.Println(today, sekarang, isToday2)

	const satu, dua = 1, 2
	const three, four string = " tiga ", "empat"

	fmt.Print(satu, dua, three, four)
}
