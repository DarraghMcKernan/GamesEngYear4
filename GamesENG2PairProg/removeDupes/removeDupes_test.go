// Author Shane, Jack & Darragh
package main

import (
	"testing"
)

func TestAllIdentical(t *testing.T) {
	var maxSizeArray [30000]int
	var oversizedArray [30001]int

	tests := []struct {
		testingArray []int
		expected     int
	}{
		{[]int{1, 1, 1, 1, 1, 1, 1, 1}, 1},         //check all duplicates
		{[]int{-100, -50, -10, -1, 0, 1, 1, 1}, 6}, //ensure negative numbers are allowed
		{[]int{1}, 1},                //smallest allowed size in spec
		{[]int{1, 2, 3, 4, 5, 6}, 6}, //make sure a normal ordered list works
		{[]int{1, 2, 2, 2, 3, 4, 4, 4, 5, 6, 7, 7, 7, 7}, 7}, //list with duplicates
		{[]int{-101, 2, 3, 4, 5, 6}, 0},                      //expected to fail - cannot be below -100 or above 100
		{[]int{1, 2, 3, 4, 5, 101}, 0},
		{[]int{1, 2, 3, 4, 3}, 0}, //must be non decreasing
		{[]int{}, 0},              //cannot be an empty array
		{maxSizeArray[:], 1},      //can be 30,000
		{oversizedArray[:], 0},    //cannot be larger than 30,000
	}

	for _, test := range tests {
		result := removeDupes(test.testingArray)
		if result != test.expected {
			t.Errorf("For input %q, expected %d but got %d", test.testingArray, test.expected, result)
		}
	}
}
