package main

import (
	"belajar-golang-dasar/helper"
	"fmt"
)

func main() {
	result := helper.SayHello("Piscki")
	fmt.Println(result)
	fmt.Println(helper.Application)
}
