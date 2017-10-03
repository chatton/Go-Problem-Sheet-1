package main

import (
	"../problem-sheets/util"
	"testing"
)

func TestReverseString(t *testing.T) {
	// create an anonymouse struct to hold the values we'll be providing and the expected outputs.
	var testCases = []struct {
		s, expected string
	}{
		{"Hello", "olleH"},
		{"", ""},
		{"World", "dlroW"},
		{"こんにちは世界", "界世はちにんこ"},
		{"mix こんにちは世界", "界世はちにんこ xim"},
		{"هِجَائِي", "يِئاَجِه"},
	}
	// iterate through all the test cases, use the util.Reverse function and make sure the result is as expected.
	for _, testCase := range testCases {
		if result := util.Reverse(testCase.s); result != testCase.expected {
			t.Errorf("String: [%s] - Actual [%s], expected [%s]", testCase.s, result, testCase.expected)
		}
	}
}

// function that tests the exported Sum function from the util package
// used in the problem sheet.
func TestSum(t *testing.T) {
	var testCases = []struct {
		numbers     []int
		expectedSum int
	}{
		{[]int{1, 2, 3}, 6},
		{[]int{4, 5, 6}, 15},
		{[]int{-10, 23, 57}, 70},
		{[]int{0, 0, 0}, 0},
		{[]int{-10, 20, 32}, 42},
	}

	for _, testCase := range testCases {
		if result := util.Sum(testCase.numbers); result != testCase.expectedSum {
			t.Errorf("Expected result %d, actual %d", testCase.expectedSum, result)
		}
	}

}
