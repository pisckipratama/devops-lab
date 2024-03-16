package main

import "fmt"

func main() {
	names := [...]string{"Piscki", "Pratama", "Refa", "Rahmah", "Reni", "Puspa", "Fenti", "Wati"}

	slice1 := names[3:5]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	slice2 := names[:4]
	fmt.Println(slice2)
	fmt.Println(len(slice2))
	fmt.Println(cap(slice2))

	slice3 := names[5:]
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	slice4 := names[:]
	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))

	fmt.Println("===========================================================")

	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}
	fmt.Println(days)

	var daySlice1 = days[5:]
	fmt.Println(daySlice1)
	daySlice1[0] = "Sabtu Baru"
	daySlice1[1] = "Minggu Baru"
	fmt.Println(daySlice1)
	fmt.Println(days)

	daySlice2 := append(daySlice1, "Libur Baru")
	daySlice2[0] = "Sabtu Lama"
	fmt.Println(daySlice1)
	fmt.Println(daySlice2)
	fmt.Println(days)

	fmt.Println("===========================================================")

	latihan := [...]string{"LNSW", "PANRB", "BPJS", "PUSINTEK", "LAPAN", "BPKH", "DENSUS 88", "KCI"}
	latihanInitSlice := latihan[:]
	latihanInitSlice[2] = "BPJS Kesehatan"
	latihanSlice := append(latihanInitSlice, "BIG", "Kominfo", "Icon+")
	fmt.Println(latihanSlice)
	fmt.Println(latihanInitSlice)
	fmt.Println(latihan)

	fmt.Println("===========================================================")

	var newSlice []string = make([]string, 2, 5)
	newSlice[0] = "Saudi Arabia"
	newSlice[1] = "Turkey"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice, "Egypt")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	newSlice2[0] = "Indonesia"
	fmt.Println(newSlice2)
	fmt.Println(newSlice)

	fmt.Println("===========================================================")

	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))
	fmt.Println(toSlice)
	copy(toSlice, fromSlice)
	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	fmt.Println("===========================================================")
	// hati-hati saat membuat array, bisa jadi ingin membuat array tapi jadinya slice atau sebaliknya
	iniArray := [...]int{1, 2, 3, 4, 5, 6}
	iniSlice := []int{1, 2, 3, 4, 5, 6}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
