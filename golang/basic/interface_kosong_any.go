package main

import "fmt"

/**
- Go-Lang bukanlah bahasa pemrograman yang berorientasi objek
- biasanya dalam pemrograman berorientasi object, ada satu data parent  di puncak yang bisa dianggap sebagai semua implementasi data yang ada di bahasa pemrograman tersebut
- Contoh di Java ada java.lang.Object
- Untuk menangani kasus seperti ini, di Go-Lang kita bisa menggunakan interface kosong
- interface kosong adalah interface yang tidak memiliki deklarasi method satupun, hal ini membuat secara otomatis semua tipe data akan menjadi implementasi nya
- Interface kosong, juga memiliki type alias bernama "any"
*/

func Ups() any {
	//return 1
	//return true
	return "Ups"
}

func main() {
	var kosong any = Ups()
	fmt.Println(kosong)
}
