package main

import (
	"fmt"
	"strings"
)

func FindSimilarData(input string, data ...string) string {

	var hasil []string

	for i := 0; i < len(data); i++ {
		wordIsTrue := strings.Contains(data[i], input)

		if wordIsTrue {
			hasil = append(hasil, data[i])
		}
	}

	hasilStr := strings.Join(hasil, ",")
	return hasilStr
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(FindSimilarData("iphone", "laptop", "iphone 13", "iphone 12", "iphone 12 pro"))
}
