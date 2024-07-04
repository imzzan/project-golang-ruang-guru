package main

import "fmt"

func BMICalculator(gender string, height int) float64 {
	if gender == "laki-laki" {
		return float64(float32((height - 100)) - (float32((height - 100)) * (float32(10) / float32(100))))
	} else if gender == "perempuan" {
		return float64(float32((height - 100)) - (float32((height - 100)) * (float32(15) / float32(100))))
	} else {
		return 0.0
	}
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(BMICalculator("perempuan", 165))
}
