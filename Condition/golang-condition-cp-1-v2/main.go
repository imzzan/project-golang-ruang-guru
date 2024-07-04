package main

import "fmt"

func GraduateStudent(score int, absent int) string {
	if score >= 70 && absent < 5 {
		return "lulus"
	} else if score < 70 || absent >= 5 {
		return "tidak lulus"
	} else {
		return ""
	}
}

// gunakan untuk melakukan debug
func main() {
	hasil := GraduateStudent(80, 5)
	fmt.Println(hasil)
}
