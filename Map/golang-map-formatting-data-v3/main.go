package main

import (
	"fmt"
	"strconv"
	"strings"
)

// TODO: answer here

func ChangeOutput(data []string) map[string][]string {

	hasil := make(map[string][]string)

	for _, data := range data {
		tokens := strings.Split(data, "-")
		label := tokens[0]
		index, _ := strconv.Atoi(tokens[1])
		firstOrLast := tokens[2]
		value := tokens[3]

		if _, ok := hasil[label]; !ok {
			hasil[label] = make([]string, 0)
		}

		if firstOrLast == "first" {
			if index >= len(hasil[label]) {
				hasil[label] = append(hasil[label], value)
			} else {
				hasil[label][index] = value + hasil[label][index]
			}
		} else {
			if index >= len(hasil[label]) {
				hasil[label] = append(hasil[label], value)
			} else {
				hasil[label][index] = hasil[label][index] + " " + value
			}
		}
	}
	fmt.Println(hasil)
	return hasil
}

// bisa digunakan untuk melakukan debug
func main() {
	data := []string{"account-0-first-John", "account-0-last-Doe", "account-1-first-Jane", "account-1-last-Doe", "address-0-first-Jaksel", "address-0-last-Jakarta", "address-1-first-Bandung", "address-1-last-Jabar", "phone-0-first-081234567890", "phone-1-first-081234567891"}
	res := ChangeOutput(data)

	fmt.Println(res)
}
