package main

import (
	"encoding/csv"
	"os"
)

func main() {
	writer := csv.NewWriter(os.Stdout)

	_ = writer.Write([]string{"piscki", "febrian", "pratama"})
	_ = writer.Write([]string{"john", "doe", "kelly"})
	_ = writer.Write([]string{"allice", "annita", "rahmah"})

	writer.Flush()
}
