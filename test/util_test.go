package main

import (
	"../problem-sheets/util"
	"testing"
)

func TestReverseString(t *testing.T) {
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
	for _, testCase := range testCases {
		if result := util.Reverse(testCase.s); result != testCase.expected {
			t.Errorf("String: [%s] - Actual [%s], expected [%s]", testCase.s, result, testCase.expected)
		}
	}
}

// function that tests the exported Sum function from the util package
// used in the problem sheet.
func TestSum(t *testing.T) {

	testCases := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{-10, 23, 57},
		{0, 0, 0},
		{-10, 20, 32},
	}

	expectedResults := []int{6, 15, 70, 0, 42}

	for index, testCase := range testCases {
		if result := util.Sum(testCase); result != expectedResults[index] {
			t.Errorf("Expected result %d, actual %d", expectedResults[index], result)
		}
	}

}
