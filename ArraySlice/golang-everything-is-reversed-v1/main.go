package main

import "strconv"

func ReverseArray(arr [5]int) [5]int {

	result := [5]int{}

	for i := 0; i < len(arr); i++ {
		result[i] = arr[len(arr)-1-i]
	}

	return result
}

func ReverseDigit(number int) int {

	s := strconv.Itoa(number)

	rev := ""

	for _, c := range s {
		rev = string(c) + rev
	}

	result, _ := strconv.Atoi(rev)

	return result

}
func ReverseData(arr [5]int) [5]int {
	numArray := ReverseArray(arr)

	for i, item := range numArray {
		numArray[i] = ReverseDigit(item)
	}

	return numArray
}
