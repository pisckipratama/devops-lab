package main

import "fmt"

func main() {
	//counter := 1

	// cara 1
	//for counter <= 10 {
	//	fmt.Println("Perulangan ke", counter)
	//	counter++
	//}

	// cara 2
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan cara 2 ke ", counter)
	}

	// for range
	names := []string{"Piscki", "Pratama", "John"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for index, name := range names {
		fmt.Println("Index", index, "=", name)
	}

	for _, name := range names {
		fmt.Println(name)
	}
}
