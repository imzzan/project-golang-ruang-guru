package main

import (
	"fmt"
	"strings"
	"unicode"
)

func ReverseWord(str string) string {
	result := ""

	kata := ""
	for i := 0; i < len(str); i++ {
		if str[i] != ' ' {
			kata += string(str[i])
		}

		if str[i] == ' ' || i == len(str)-1 {
			reverse := ""

			for j := len(kata) - 1; j >= 0; j-- {
				reverse += string(kata[j])
			}

			if unicode.IsUpper(rune(kata[0])) == true {
				reverse = strings.ToUpper(string(reverse[0])) + reverse[1:]
			}

			if unicode.IsLower(rune(kata[len(kata)-1])) == true {
				reverse = reverse[:len(reverse)-1] + strings.ToLower(string(reverse[len(reverse)-1]))
			}
			result += reverse + " "
			kata = ""
		}
	}

	hasil := result[:len(result)-1]

	return hasil

}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(ReverseWord("Aku Sayang Ibu"))
	fmt.Println(ReverseWord("A bird fly to the Sky"))
}
