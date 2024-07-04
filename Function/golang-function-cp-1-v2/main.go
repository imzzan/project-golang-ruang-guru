package main

import (
	"fmt"
)

func DateFormat(day, month, year int) string {
	var dayFmt string
	monthFmt := map[int]string{
		1:  "January",
		2:  "February",
		3:  "March",
		4:  "April",
		5:  "May",
		6:  "June",
		7:  "July",
		8:  "August",
		9:  "September",
		10: "October",
		11: "November",
		12: "December",
	}

	for i := 1; i <= day; i++ {
		if i < 10 {
			dayFmt = fmt.Sprintf("0%d", i)
		} else {
			dayFmt = fmt.Sprintf("%d", i)
		}
	}

	return fmt.Sprintf("%s-%s-%d", dayFmt, monthFmt[month], year)
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(DateFormat(1, 1, 2012))
}
