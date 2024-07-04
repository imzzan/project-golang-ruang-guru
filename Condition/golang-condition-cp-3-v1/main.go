package main

import "fmt"

func GetPredicate(math, science, english, indonesia int) string {
	average := (math + science + english + indonesia) / 4

	if average == 100 {
		return "Sempurna"
	} else if average >= 90 {
		return "Sangat Baik"
	} else if average >= 80 {
		return "Baik"
	} else if average >= 70 {
		return "Cukup"
	} else if average >= 60 {
		return "Kurang"
	} else {
		return "Sangat kurang"
	}
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(GetPredicate(50, 80, 100, 60))
}
