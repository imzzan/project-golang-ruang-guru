package main

import (
	"fmt"
	"strconv"
)

func BiggestPairNumber(numbers int) int {
	numberStr := strconv.Itoa(numbers)

	var terbesar = 0
	var pair = 0

	for i := 1; i < len(numberStr); i++ {
		number1, _ := strconv.Atoi(string(numberStr[i]))
		number2, _ := strconv.Atoi(string(numberStr[i-1]))

		if number1+number2 > terbesar {
			terbesar = number1 + number2

			pairNumb, _ := strconv.Atoi(string(numberStr[i-1]) + string(numberStr[i]))
			pair = pairNumb
		}
	}

	return pair
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(BiggestPairNumber(11223344))
}
