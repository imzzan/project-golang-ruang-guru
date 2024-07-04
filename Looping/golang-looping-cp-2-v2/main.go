package main

import (
	"fmt"
	"strings"
)

// hello World => d_l_r_o_W o_l_l_e_H
func ReverseString(str string) string {
	var hasil []string
	for i := 0; i < len(str); i++ {
		hasil = strings.Split(str, "")
		for i2, j := 0, len(hasil)-1; i2 < j; i2, j = i2+1, j-1 {
			hasil[i2], hasil[j] = hasil[j], hasil[i2]
		}

	}
	hasilStr := strings.Join(hasil, "")
	hasil = strings.Split(hasilStr, " ")
	for k, word := range hasil {
		reversedWord := ""
		for l := 0; l < len(word); l++ {
			reversedWord += string(word[l]) + "_"
		}
		reversedWord = strings.TrimSuffix(reversedWord, "_")
		hasil[k] = reversedWord
	}
	return strings.Join(hasil, " ")
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(ReverseString("I am a student"))
}
