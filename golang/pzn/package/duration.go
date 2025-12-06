package main

import (
	"fmt"
	"time"
)

func main() {
	var duration1 time.Duration = 100 * time.Second
	var duration2 time.Duration = 10 * time.Millisecond
	var duration3 time.Duration = duration1 - duration2

	fmt.Printf("duration 1: %d\n", duration1)
	fmt.Printf("duration 2: %d\n", duration2)
	fmt.Printf("duration total: %d\n", duration3)
}
