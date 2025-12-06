package main

import "fmt"

/**
- sebelumnya untuk membuat pointer dengan menggunakan operator &
- golang juga memiliki function new yang bisa digunakan untuk membuat pointer
- namun function new hanya mengembalikan pointer ke data kosong, artinya tidak ada data awal
*/

type Address struct {
	City, Province, Country string
}

func main() {
	var alamat1 *Address = new(Address)
	var alamat2 *Address = alamat1

	alamat2.Country = "Indonesia"
	alamat1.City = "Sumedang"
	alamat2.Province = "West Java"

	fmt.Println(alamat1)
	fmt.Println(alamat2)
}
