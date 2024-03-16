package main

import "fmt"

/**
- struct adalah sebuah template data yang digunakan untuk menggabungkan nol atau lebih tipe data lainnya dalam satu kesatuan
- struct biasanya representasi data dalam program aplikasi yang kita buat
- data di struct disimpan dalam field
- sederhananya struct adalah kumpulan dari field
*/

/**
- struct merupakan template data atau prototype data
- struct tidak bisa langsung digunakan
- namaun kita bisa membuat data/object dari struct yang telah kita buat
*/

/**
==== Struct Method ====
- struct adalah tipe data seperti tipe data lainnya, dia bisa digunakan sebagai parameter untuk function
- naman jika ingin menambahkan method ke dalam structs, sehingga seakan-akan sebuah struct memiliki function
- method adalah function
*/

type Customer struct {
	Name, Address string
	Age           int
}

// contoh koding struct method
func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "my name is", customer.Name)
}

func main() {
	var piscki Customer

	piscki.Name = "Piscki Pratama"
	piscki.Address = "Jakarta"
	piscki.Age = 24

	fmt.Println(piscki)
	fmt.Println(piscki.Age)

	/**
	=== Struct Literals ===
	*/

	refa := Customer{
		Name:    "Refa Rahmah",
		Age:     25,
		Address: "Bandung",
	}
	fmt.Println(refa)
	fmt.Println(refa.Name)

	hanin := Customer{"Haninda Septrianiavi", "Bandung", 25}
	fmt.Println(hanin)
	fmt.Println(hanin.Name)

	// memanggil struct method
	hanin.sayHello("Piscki")
}
