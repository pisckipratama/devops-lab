package main

import "fmt"

func main() {
	// 13.1 belajar basic if-else
	var point = 10

	if point == 10 {
		fmt.Println("lulus dengan nilai sempurna")
	} else if point >= 8 {
		fmt.Println("lulus dengan nilai memuaskan")
	} else if point >= 5 {
		fmt.Println("lulus")
	} else if point == 4 {
		fmt.Println("hampir lulus")
	} else {
		fmt.Printf("tidak lulus. nilai anda %d\n", point)
	}

	// 13.2 belajar varible temporary pada if-else
	var point1 = 6840.0

	if percent := point1 / 100; percent >= 100 {
		fmt.Printf("%.1f%s perfect!\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s good!\n", percent, "%")

	} else {
		fmt.Printf("%.1f%s not bad\n", percent, "%")
	}

	// Deklarasi variabel temporary hanya bisa dilakukan lewat metode type inference yang menggunakan tanda :=.
	// Penggunaan keyword var di situ tidak diperbolehkan karena menyebabkan error.

	// 13.3 belajar switch case
	var point2 = 8

	switch point2 {
	case 8:
		fmt.Println("Perfect!")
	case 7:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}

	// Perlu diketahui, switch pada pemrograman Go memiliki perbedaan dibanding bahasa lain.
	// Di Go, ketika sebuah case terpenuhi, tidak akan dilanjutkan ke pengecekan case selanjutnya, meskipun tidak ada keyword break di situ.
	// Konsep ini berkebalikan dengan switch pada umumnya pemrograman lain (yang ketika sebuah case terpenuhi,
	// maka akan tetap dilanjut mengecek case selanjutnya kecuali ada keyword break).

	// 13.4 case untuk banyak kondisi
	var point3 = 3

	switch point3 {
	case 9, 8:
		fmt.Println("Perfect!")
	case 7, 6, 5, 4:
		fmt.Println("awesome")
	case 3:
		fmt.Println("nice try")
	default:
		fmt.Println("try again")
	}

	// 13.5 kurung kurawal pada keyword case dan default
	switch point3 {
	case 8:
		fmt.Println("Perfect!")
	case 7, 6, 5, 4:
		fmt.Println("awesome")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you can be better!")
		}
	}

	// 13.6 belajar switch case dengan if-else style
	var point4 = 2

	switch {
	case point4 == 8:
		fmt.Println("Perfect")
	case (point4 < 8) && (point4 > 3):
		fmt.Println("awesome")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("You need to learn more")
		}
	}

	// 13.7 belajar menggunakan fallthrough dalam switch
	var point5 = 7

	switch {
	case point5 == 8:
		fmt.Println("perfect broo!")
	case (point5 < 8) && (point5 > 6):
		fmt.Println("awesome mamen!")
		fallthrough
	case point5 < 6:
		fmt.Println("nice try, you need to learn more")
	default:
		{
			fmt.Println("Not bad")
			fmt.Println("you need to focus learning")
		}
	}

	// 13.8 belajar menggunakan if-else bersarang
	var angka = 10
	if angka > 7 {
		switch angka {
		case 10:
			fmt.Println("Perfect, you almost be the champion")
		default:
			fmt.Println("ganbatte!")
		}
	} else {
		if angka == 5 {
			fmt.Println("come on, more fighting")
		} else if angka == 3 {
			fmt.Println("keep practice")
		} else {
			fmt.Println("you can do it")
			if angka == 0 {
				fmt.Println("try harder!")
			}
		}
	}
}
