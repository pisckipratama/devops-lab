package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	/* Tidak Menggunakan Pointer */
	var address1 Address = Address{"Bandung", "Jawa Barat", "Indonesia"}
	var address2 *Address = &address1
	address2.City = "Depok"
	fmt.Println(address1)
	fmt.Println(address2)

	// address2 = &Address{"Jakarta Pusat", "DKI Jakarta", "Indonesia"}
	*address2 = Address{"Jakarta Pusat", "DKI Jakarta", "Indonesia"}
	fmt.Println(address1)
	fmt.Println(address2)
}
