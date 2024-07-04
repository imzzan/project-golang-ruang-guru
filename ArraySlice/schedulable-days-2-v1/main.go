package main

import "fmt"

func SchedulableDays(villager [][]int) []int {
	availableDate := map[int]int{}

	for _, orang := range villager {
		for _, tanggal := range orang {
			availableDate[tanggal] = availableDate[tanggal] + 1
		}
	}

	hasil := []int{}
	for key, value := range availableDate {
		if value == len(villager) {
			hasil = append(hasil, key)
		}
	}

	return hasil
}

func main() {
	fmt.Println(SchedulableDays([][]int{
		{7, 12, 19, 22},
		{12, 19, 21, 23},
		{7, 12, 19},
		{12, 19},
	}))
}
