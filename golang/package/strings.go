package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Piscki Pratama", "Piscki"))
	fmt.Println(strings.Split("Piscki Pratama", " "))
	fmt.Println(strings.ToLower("Piscki Pratama"))
	fmt.Println(strings.ToUpper("Piscki Pratama"))
	fmt.Println(strings.ReplaceAll("Piscki Pratama Piscki Feb", "Piscki", "Tama"))
	fmt.Println(strings.Trim("         Piscki Pratama        ", " "))

}
