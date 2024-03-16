package main

import "fmt"

/**
- tipe assertions merupakan kemampuan merubah tipe data menjadi tipe data yang diinginkan
- fitur ini sering kali digunakan ketika kita bertemu dengan data interface kosong
*/

func random() interface{} {
	return false
}

func main() {
	result := random()
	//resultString := result.(string)
	//
	//fmt.Println(resultString)
	//
	//resultInt := result.(int) // ini bakal panic
	//fmt.Println(resultInt)

	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	default:
		fmt.Println("Unknown", value)
	}
}
