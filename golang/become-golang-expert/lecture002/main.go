package main

import "fmt"

func main() {
	nums := []int{}
	fmt.Println(len(nums), cap(nums))

	nums = append(nums, 10)
	fmt.Println(len(nums), cap(nums))

	nums = append(nums, 20)
	fmt.Println(len(nums), cap(nums))

	nums = append(nums, 30)
	fmt.Println(len(nums), cap(nums))

	cities := []string{"Bandung", "Jakarta", "Tangerang"}
	fmt.Println(len(cities), cap(cities))

	provinces := make([]int, 38, 40)
	fmt.Println(len(provinces), cap(provinces))

	// copy slice
	src := []int{1, 2, 3}
	dst := make([]int, len(src))

	copy(dst, src)
	fmt.Println(dst)

	// map
	user := map[string]string{
		"name": "Andi",
		"city": "Jakarta",
	}

	user["age"] = "25"
	user["name"] = "Piscki"

	fmt.Println(user["name"])
	fmt.Println(user["city"])
	fmt.Println(user)

	value, exists := user["isMarried"]

	if exists {
		fmt.Println("isMarried: ", value)
	} else {
		fmt.Println("isMarried not found")
	}

	delete(user, "city")
	fmt.Println(user)

	for key, value := range user {
		fmt.Println(key, "=", value)
	}
}
