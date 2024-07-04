package main

import (
	"fmt"
	"strings"
)

func SlurredTalk(words *string) {
	*words = strings.Map(func(r rune) rune {
		if r == 'S' || r == 'R' || r == 'Z' {
			return 'L'
		} else if r == 'r' || r == 's' || r == 'z' {
			return 'l'
		} else {
			return r
		}
	}, *words)
}

func main() {
	// bisa dicoba untuk pengujian test case
	var words string = "Saya Steven"
	SlurredTalk(&words)
	fmt.Println(words)
}
