package main

func SchedulableDays(date1 []int, date2 []int) []int {
	result := make([]int, 0)

	for i := 1; i <= 31; i++ {
		isImamAvailable := false
		isPermanaAvailable := false

		for _, date := range date1 {
			if date == i {
				isImamAvailable = true
			}
		}

		for _, date := range date2 {
			if date == i {
				isPermanaAvailable = true
			}
		}

		if isImamAvailable && isPermanaAvailable {
			result = append(result, i)
		}
	}

	return result
}

func main() {
	SchedulableDays([]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5})
}
