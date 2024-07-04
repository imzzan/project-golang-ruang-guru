package main

import (
	"fmt"
	"sort"
)

func Sortheight(height []int) []int {
	sort.Slice(height, func(i, j int) bool {
		return height[i] < height[j]
	})

	return height
}

func main() {
	fmt.Println(Sortheight([]int{19,12,1}))
}
