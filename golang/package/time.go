package main

import (
	"fmt"
	"time"
)

func main() {
	var now time.Time = time.Now()

	fmt.Println(now)

	var utc time.Time = time.Date(2000, time.February, 24, 1, 0, 0, 0, time.UTC)
	fmt.Println(utc)
	fmt.Println(utc.Local())

	formatter := "2006-01-02 15:04:05"
	value := "2022-01-31 10:10:00"
	valueTime, err := time.Parse(formatter, value)

	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(valueTime)
	}

	fmt.Println(valueTime.Year())
	fmt.Println(valueTime.Month())
	fmt.Println(valueTime.Day())
	fmt.Println(valueTime.Hour())
	fmt.Println(valueTime.Second())

	parse, err := time.Parse(time.RFC3339, "2022-09-29T17:39:11Z")
	if err != nil {
		fmt.Print("Error", err.Error())
	} else {
		fmt.Println(parse)
	}
}
