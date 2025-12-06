package main

import (
	"fmt"
)

func main() {
	// 14.1 belajar for dasar
	for i := 0; i < 5; i++ {
		fmt.Printf("Angka ke-%d\n", i)
	}

	fmt.Println("=========== end of 14.1 ===========")

	// 14.2 menggunakan keyword for dengan argument hanya kondisi
	var number = 0
	for number < 5 {
		fmt.Println("angka", number)
		number++
	}

	fmt.Println("=========== end of 14.2 ===========")

	// 14.3 menggunakan keyword for tanpa argument
	var number1 = 0
	for {
		fmt.Println("Angka", number1)

		number1++
		if number1 == 7 {
			break
		}
	}

	fmt.Println("=========== end of 14.3 ===========")

	// 14.4 penggunaan keyword for - range
	var xs = "123" // string
	for i, v := range xs {
		fmt.Println("Index=", i, "Value=", v)
	}

	var ys = [5]int{10, 20, 30, 40, 50} // array
	for _, v := range ys {
		fmt.Println("Value=", v)
	}

	var zs = ys[0:2] // slice
	for _, v := range zs {
		fmt.Println("Value=", v)
	}

	var kvs = map[byte]int{'a': 0, 'b': 1, 'c': 2} // map
	for k, v := range kvs {
		fmt.Println("Key=", k, "Value=", v)
	}

	// boleh juga baik k dan v nya diabaikan
	for range kvs {
		fmt.Println("done")
	}

	// selain itu bisa juga dengan cukup menentukan nilai numerik perulangan
	for i := range 5 {
		fmt.Print(i)
	}

	fmt.Println("")
	fmt.Println("=========== end of 14.4 ===========")

	// 14.5 penggunaan keyword break dan continue
	for i := 1; i <= 10; i++ {
		if i%2 == 1 {
			continue
		}

		if i > 8 {
			break
		}

		fmt.Println("Angka", i)
	}

	fmt.Println("=========== end of 14.5 ===========")

	// 14.6 belajar for bersarang
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print(j, "")
		}

		fmt.Println()
	}

	fmt.Println("=========== end of 14.6 ===========")

	// 14.7 belajar label dalam for
outerLoop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 3 {
				break outerLoop
			}
			fmt.Print("matriks [", i, "][", j, "]", "\n")
		}
	}
}
