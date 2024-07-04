package main

import "fmt"

func CountProfit(data [][][2]int) []int {
	result := map[int]int{}
	for _, item := range data {
		for i, bulan := range item {
			income := bulan[0]
			expanse := bulan[1]

			profit := income - expanse
			result[i+1] += profit
		}
	}

	jumlahBulan := 0
	for k := range result {
		if k > jumlahBulan {
			jumlahBulan = k
		}
	}

	fmt.Println(jumlahBulan)
	hasil := make([]int, jumlahBulan)

	for key, value := range result {
		hasil[key-1] = value
	}

	return hasil
}

func main() {
	fmt.Println(CountProfit([][][2]int{
		{{1000, 500}, {500, 200}},
		{{1200, 200}, {1000, 800}},
		{{500, 100}, {700, 100}},
	}))
}
