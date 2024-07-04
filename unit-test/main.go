package main

// Instruction: Implement a function IsValidEmail that checks if a given string is a valid email address. For simplicity, assume an email is valid if it contains exactly one "@" symbol. Test this function.

import "strings"

func IsValidEmail(email string) bool {
	return strings.Count(email, "@") == 1
}

func main() {
	IsValidEmail("muzani@gmail.com")
}
