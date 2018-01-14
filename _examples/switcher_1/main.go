package main

import (
	"os"

	goStackQuery "github.com/dreddsa5dies/goStackQuery"
)

func main() {
	file, err := os.Open("file")
	if err != nil {
		goStackQuery.Query(err, 1)
	}
	defer file.Close()

	file.WriteString("Aloha")
}
