package main

import (
	"fmt"
)

func FindMin(nums ...int) int {
	min := nums[0]
	for _, item := range nums {
		if item < min {
			min = item
		}
	}
	fmt.Println(min)
	return min
}

func FindMax(nums ...int) int {
	max := nums[0]
	for _, item := range nums {
		if item > max {
			max = item
		}
	}
	return max
}

func SumMinMax(nums ...int) int {
	minAngka := FindMin(nums...)
	maxAngka := FindMax(nums...)

	return minAngka + maxAngka
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(SumMinMax(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
}
