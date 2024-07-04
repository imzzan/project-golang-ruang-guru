package main

import (
	"fmt"
	"strings"
)

func EmailInfo(email string) string {
	emailArray := strings.Split(email, "@")[1]
	emailArray2 := strings.Split(emailArray, ".")
	domain := ""
	tld := ""

	for _, word := range emailArray {
		if word != '.' {
			domain = string(emailArray2[0])
		} else {
			tld = string(emailArray2[1])
		}
	}

	return fmt.Sprintf("Domain: %s dan TLD: %s", domain, tld)
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(EmailInfo("admin@yahoo.com"))
}
