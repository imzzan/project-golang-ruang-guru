package main

import "testing"

func TestIsValidEmail(t *testing.T) {
	email := "muzani@gmail.com"
	if !IsValidEmail(email) {
		t.Errorf("Expected %s valid email is", email)
	}
}

func TestInvalidEmail(t *testing.T) {
	email := "muzani&gmail.com"
	if IsValidEmail(email) {
		t.Errorf("Expected %s invalid email is", email)
	}
}

func TestNoAtSymbol(t *testing.T) {
	email := "muzani.com"
	if IsValidEmail(email) {
		t.Errorf("Expected %s invalid email", email)
	}
}
