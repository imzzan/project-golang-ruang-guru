package main

import (
	"fmt"
	"strings"
)

func CountingLetter(text string) int {
	counter := 0
	for _, letter := range text {
		capital := strings.ToUpper(string(letter))
		if capital == "R" || capital == "S" || capital == "T" || capital == "Z" {
			counter++
		}
	}
	return counter
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(CountingLetter("Semangat"))
}
