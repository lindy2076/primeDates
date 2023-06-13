package main

import "testing"

// calls IsPrime from prime.go, checks if it works properly
func TestIsPrime(t *testing.T) {
	var data = []struct {
		name     string
		input    uint32
		expected bool
	}{
		{"0 is not prime", 0, false},
		{"1 is not prime", 1, false},
		{"2 is prime", 2, true},
		{"4 is not prime", 4, false},
		{"9 is not prime", 9, false},
		{"17 is prime", 17, true},
		{"289 is not prime", 289, false},
	}

	for _, tt := range data {
		t.Run(tt.name, func(t *testing.T) {
			res := IsPrime(tt.input)
			if res != tt.expected {
				t.Errorf("IsPrime (%d) got %t, expected %t", tt.input, res, tt.expected)
			}
		})
	}
}
