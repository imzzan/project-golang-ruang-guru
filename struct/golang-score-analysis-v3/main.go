package main

import (
	"fmt"
	"math"
)

type School struct {
	Name    string
	Address string
	Grades  []int
}

func (s *School) AddGrade(grades ...int) {
	s.Grades = append(s.Grades, grades...)
}

func Analysis(s School) (float64, int, int) {
	if len(s.Grades) == 0 {
		return 0, 0, 0
	}
	avg := 0.0
	min := s.Grades[0]
	max := s.Grades[0]

	for i := 0; i < len(s.Grades); i++ {
		avg += float64(s.Grades[i])
		if s.Grades[i] < min {
			min = s.Grades[i]
		}
		if s.Grades[i] > max {
			max = s.Grades[i]
		}
	}
	if len(s.Grades) != 0 {
		avg /= float64(len(s.Grades))
		avg = math.Round(avg*100) / 100
	}
	return avg, min, max
}

// gunakan untuk melakukan debugging
func main() {
	avg, min, max := Analysis(School{
		Name:    "Imam Assidiqi School",
		Address: "Jl. Imam Assidiqi",
		Grades:  []int{100, 90, 80, 70, 60},
	})

	fmt.Println(avg, min, max)
}
