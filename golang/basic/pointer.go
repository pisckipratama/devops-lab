package main

import "fmt"

/**
- secara default di Go semua varible itu dipassing by value, bukan by reference
- artinya, jika kita mengirim sebuah variable ke dalam function, method atau variable lain, sebenernya yang dikirim adalah duplikasi valuenya.
*/

/**
- pinter adalah kemampuan membuat references ke lokasi data di memory yang sama, tanpa menduplikasi data yang sudah ada
- sederhananya, dengan kemampuan pointer, kita bisa membuat pass by reference
*/

/**
Untuk membuat sebuah variable dengan nilai pointer ke variable yang lain, kita bisa menggunakan operator '&' diikuti dengan nama variable nya
*/

type Address struct {
	City, Province, Country string
}

func main() {
	/* Tidak Menggunakan Pointer */
	address1 := Address{"Bandung", "Jawa Barat", "Indonesia"}
	address2 := address1

	address2.City = "Depok"

	fmt.Println(address1)
	fmt.Println(address2)

	/* Mmenggunakan Pointer cara-1 */
	address3 := Address{"Surabaya", "Jawa Timur", "Indonesia"}
	address4 := &address3

	address4.City = "Solo"
	address4.Province = "Jawa Tengah"

	fmt.Println(address3)
	fmt.Println(address4)
	fmt.Println("ini nama provinsi dari address3", address3.Province)
	fmt.Println("ini nama provinsi dari address4", address4.Province)

	/* Menggunakan Pointer cara 2*/
	var address5 Address = Address{"Gianyar", "Bali", "Indonesia"}
	var address6 *Address = &address5

	address6.City = "Balikpapan"
	address6.Province = "Kalimantan Timur"

	fmt.Println(address5)
	fmt.Println(address6)
}
