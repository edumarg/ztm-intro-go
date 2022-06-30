package main

import (
	"fmt"
	"testing"
)

func TestIsValidEmail(t *testing.T) {
	data := "email@example.com"
	if !IsValidEmail(data) {
		t.Errorf("IsValidEmail(%v)=false, want true", data)
	}
}

// test table
func TestIsValidEmailTable(t *testing.T) {
	fmt.Println("Test table")
	table := []struct {
		email string
		want  bool
	}{
		{"email@example.com", true},
		{"email@example", false},
		{"email", false},
	}

	for _, data := range table {
		result := IsValidEmail(data.email)
		if result != data.want {
			t.Errorf("%v: %t, want %t", data.email, result, data.want)
		}
	}

}
