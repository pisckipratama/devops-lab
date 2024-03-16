package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

func main() {
	fmt.Println(sumAll(10, 10, 10))
	fmt.Println(sumAll(10, 10, 10, 10, 10, 10, 5))
	fmt.Println(sumAll(10, 10, 10, 10, 10, 10, 10, 5, 5, 10))

	numbers := []int{20, 20, 20, 15, 10, 5}

	fmt.Println(sumAll(numbers...))
}
