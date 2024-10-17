// Authors
// Shane Conroy , Jack Nulty and Darragh McKernan
package main

import (
	"testing"
)

func TestGenerateNumbersFromString(t *testing.T) {
	tests := []struct {
		roman    string
		expected int
	}{
		// Simple cases
		{"I", 1},    // Single numeral
		{"V", 5},    // Single numeral
		{"X", 10},   // Single numeral
		{"L", 50},   // Single numeral
		{"C", 100},  // Single numeral
		{"D", 500},  // Single numeral
		{"M", 1000}, // Single numeral

		// Addition cases
		{"II", 2},     // 1 + 1 = 2
		{"III", 3},    // 1 + 1 + 1 = 3
		{"VI", 6},     // 5 + 1 = 6
		{"VII", 7},    // 5 + 1 + 1 = 7
		{"VIII", 8},   // 5 + 1 + 1 + 1 = 8
		{"XXVII", 27}, // 10 + 10 + 5 + 1 + 1 = 27
		{"XII", 12},   // 10 + 1 + 1 = 12
		{"XXX", 30},   // 10 + 10 + 10 = 30

		// Subtraction cases
		{"IV", 4},   // 5 - 1 = 4
		{"IX", 9},   // 10 - 1 = 9
		{"XL", 40},  // 50 - 10 = 40
		{"XC", 90},  // 100 - 10 = 90
		{"CD", 400}, // 500 - 100 = 400
		{"CM", 900}, // 1000 - 100 = 900

		// Complex cases
		{"XIV", 14},         // 10 + 4 = 14
		{"XCIV", 94},        // 90 + 4 = 94
		{"MCMXCIV", 1994},   // 1000 + 900 + 90 + 4 = 1994
		{"MMXX", 2020},      // 1000 + 1000 + 10 + 10 = 2020
		{"MMMCMXCIX", 3999}, // 1000 + 1000 + 1000 + 900 + 90 + 9 = 3999

		// Edge cases
		{"I", 1},            // Minimum
		{"MMMCMXCIX", 3999}, // Maximum
	}

	for _, test := range tests {
		result := generateNumbersFromString(test.roman)
		if result != test.expected {
			t.Errorf("For input %q, expected %d but got %d", test.roman, test.expected, result)
		}
	}
}
