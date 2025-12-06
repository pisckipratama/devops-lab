package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func (u User) SayHello() {
	fmt.Println("Hello,", u.Name)
}

func changeValue(x *int) {
	*x = 100
}

func add(a int, b int) int {
	return a + b
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

func main() {
	var name string = "Piscki"
	age := 25
	height := 167.8

	fmt.Println(name, age, height)

	fmt.Println("Hello from Go!")

	result := add(10, 90)
	fmt.Println(result)

	u := User{Name: "Pratama", Age: 25}
	fmt.Println(u.Name)
	fmt.Println(u.Age)
	u.SayHello()

	a := 10
	changeValue(&a)
	fmt.Println(a)

	res, err := divide(10, 5)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println("Result: ", res)
}
