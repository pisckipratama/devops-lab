package main

import (
	"fmt"
	"slices"
)

func main() {
	names := []string{"John", "Ringo", "Doe", "Tom", "Jenny", "Rahmah"}
	value := []int{10, 90, 20, 225, 115, 205, 100, 5, 20}

	fmt.Println(slices.Min(names))
	fmt.Println(slices.Min(value))
	fmt.Println(slices.Max(names))
	fmt.Println(slices.Max(value))
	fmt.Println(slices.Contains(names, "Piscki"))
	fmt.Println(slices.Contains(names, "Rahmah"))
	fmt.Println(slices.Index(names, "Tom"))
	fmt.Println(slices.Index(names, "Jelly"))

}
