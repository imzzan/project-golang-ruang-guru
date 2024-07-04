package main

import "fmt"

func CountingNumber(n int) float64 {
	var hasil float64 = 0
	var i float64
	for i = 1.0; i <= float64(n); i += 0.5 {
		hasil += i
	}
	return hasil
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(CountingNumber(10))
}
