package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex = regexp.MustCompile(`e([a-z])o`)

	fmt.Println(regex.MatchString("eko"))
	fmt.Println(regex.MatchString("edo"))
	fmt.Println(regex.MatchString("eDo"))

	fmt.Println(regex.FindAllString("eko edo egi ego e1o e2o eka abo eto eno aja", 5))
}
