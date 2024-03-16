package main

import "fmt"

/**
- walaupun method akan menempel di struct, tapi sebenernya data struct yang diakses di dalam method adalah pass by value
- sangat direkomendasikan menggunakan pointer di method, sehigga tidak boros memory karena harus selalu diduplikasi ketika memanggil method
*/

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	piscki := Man{"Piscki"}
	piscki.Married()

	fmt.Println(piscki.Name)
}
