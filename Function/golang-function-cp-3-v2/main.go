package main

import (
	"fmt"
	"strings"
)

func FindShortestName(names string) string {

	var strArray []string

	for i := 0; i < len(names); i++ {
		if string(names[i]) == " " {
			fmt.Println(string(names[i]))
			strArray = strings.Split(names, " ")
			break
		} else if string(names[i]) == "," {
			strArray = strings.Split(names, ",")
			break
		} else if string(names[i]) == ";" {
			strArray = strings.Split(names, ";")
			break
		}
	}

	shortName := strArray[0]
	shortLength := len(shortName)
	for _, name := range strArray[1:] {
		if len(name) < shortLength {
			shortName = name
			shortLength = len(name)
		}
	}

	return shortName
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(FindShortestName("Hanif Joko Tio Andi Budi Caca Hamdan")) // "Tio"
	fmt.Println(FindShortestName("Budi;Tia;Tio"))                         // "Tia"
}
