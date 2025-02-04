package utils_test

import (
	"number-classifier/utils"
	"testing"
)

func TestDigitSum(t *testing.T) {
	tests := []struct {
		name     string
		number   int
		expected int
	}{
		{"single digit", 5, 5},
		{"two digits", 23, 5},
		{"three digits", 123, 6},
		{"negative number", -123, 6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := utils.DigitSum(tt.number)
			if result != tt.expected {
				t.Errorf("DigitSum(%d) = %d; want %d", tt.number, result, tt.expected)
			}
		})
	}
}

func TestIsArmstrong(t *testing.T) {
	tests := []struct {
		name     string
		number   int
		expected bool
	}{
		{"armstrong number", 153, true},
		{"not armstrong number", 123, false},
		{"negative number", -153, false},
		{"single digit", 5, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := utils.IsArmstrong(tt.number)
			if result != tt.expected {
				t.Errorf("IsArmstrong(%d) = %t; want %t", tt.number, result, tt.expected)
			}
		})
	}
}

func TestIsPerfect(t *testing.T) {
	tests := []struct {
		name     string
		number   int
		expected bool
	}{
		{"perfect number", 28, true},
		{"not perfect number", 123, false},
		{"negative number", -28, false},
		{"single digit", 5, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := utils.IsPerfect(tt.number)
			if result != tt.expected {
				t.Errorf("IsPerfect(%d) = %t; want %t", tt.number, result, tt.expected)
			}
		})
	}
}

func TestIsPrime(t *testing.T) {
	tests := []struct {
		name     string
		number   int
		expected bool
	}{
		{"prime number", 7, true},
		{"not prime number", 9, false},
		{"negative number", -7, false},
		{"single digit", 5, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := utils.IsPrime(tt.number)
			if result != tt.expected {
				t.Errorf("IsPrime(%d) = %t; want %t", tt.number, result, tt.expected)
			}
		})
	}
}
