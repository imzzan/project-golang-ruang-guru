package main

import (
	"fmt"
	"strings"
)

func CountVowelConsonant(str string) (int, int, bool) {
	str = strings.ToLower(str)
	arrayStr := strings.Split(str, " ")
	str = strings.Join(arrayStr, "")
	arrayStr = strings.Split(str, "'")
	str = strings.Join(arrayStr, "")
	arrayStr = strings.Split(str, ",")
	str = strings.Join(arrayStr, "")
	jumlahHurufVokal := 0
	jumlahHurufKonsonan := 0
	var chaeck bool

	for i := 0; i < len(str); i++ {
		if string(str[i]) == "a" || string(str[i]) == "i" || string(str[i]) == "u" || string(str[i]) == "e" || string(str[i]) == "o" {
			jumlahHurufVokal += 1
			chaeck = false
		} else {
			jumlahHurufKonsonan += 1
			chaeck = false
		}
	}

	if jumlahHurufKonsonan == 0 || jumlahHurufVokal == 0 {
		chaeck = true
	}

	return jumlahHurufVokal, jumlahHurufKonsonan, chaeck
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(CountVowelConsonant("bbbbb ccccc"))
}
