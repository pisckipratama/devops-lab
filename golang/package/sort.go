package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type UserSlice []User

func (s UserSlice) Len() int {
	return len(s)
}

func (s UserSlice) Less(i, j int) bool {
	return s[i].Age < s[j].Age
}

func (s UserSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	users := []User{
		{"Piscki", 24},
		{"John", 27},
		{"Hernandez", 22},
		{"Gilbert", 26},
		{"Noah", 28},
	}

	sort.Sort(UserSlice(users))
	fmt.Println(users)
}
