// Authors
// Shane Conroy , Jack Nulty and Darragh McKernan
// Pair programed in brain waves whoever had an idea took the wheel (Jesus wasnt available)
// Description
/*
	Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.
	Symbol       Value
	I             1
	V             5
	X             10
	L             50
	C             100
	D             500
	M             1000
	For example, 2 is written as II in Roman numeral, just two ones added together. 12 is written as XII, which is simply X + II. The number 27 is written as XXVII, which is XX + V + II.
	Roman numerals are usually written largest to smallest from left to right. However, the numeral for four is not IIII. Instead, the number four is written as IV. Because the one is before the five we subtract it making four. The same principle applies to the number nine, which is written as IX. There are six instances where subtraction is used:
	I can be placed before V (5) and X (10) to make 4 and 9.
	X can be placed before L (50) and C (100) to make 40 and 90.
	C can be placed before D (500) and M (1000) to make 400 and 900.
	Given a roman numeral, convert it to an integer.

	Example 1:
	Input: s = "III"
	Output: 3
	Explanation: III = 3.
	Example 2:
	Input: s = "LVIII"
	Output: 58
	Explanation: L = 50, V= 5, III = 3.
	Example 3:
	Input: s = "MCMXCIV"
	Output: 1994
	Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.

	Constraints:
	1 <= s.length <= 15
	s contains only the characters ('I', 'V', 'X', 'L', 'C', 'D', 'M').
	It is guaranteed that s is a valid roman numeral in the range [1, 3999].
*/
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
