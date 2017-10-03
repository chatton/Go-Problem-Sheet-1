package main

import (
	"../problem-sheets/util" // the package under test
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

func TestAsIntSlice(t *testing.T) {
	var testCases = []struct {
		line   string
		result []int
	}{
		{"1 2 3 4 5", []int{1, 2, 3, 4, 5}},
		{"2 34 5 6", []int{2, 34, 5, 6}},
		{"-2 32 -43 23", []int{-2, 32, -43, 23}},
		{"-3 43 23 12 0 0 0", []int{-3, 43, 23, 12, 0, 0, 0}},
		{"", []int{}},
	}

	for _, testCase := range testCases {
		if result, err := util.AsIntSlice(testCase.line); !compare(result, testCase.result) {
			t.Errorf("Provided line %s, expected result was %v, actual result was %v. Error %s",
				testCase.line, testCase.result, result, err)
		}
	}
}

// helper function to compare equality of 2 int slices
func compare(slice1, slice2 []int) bool {
	// can only compare slices to nil, not each other
	if slice1 == nil || slice2 == nil {
		return false
	}

	if len(slice1) != len(slice2) {
		return false
	}

	// manually compare every element.
	for index, _ := range slice1 {
		if slice1[index] != slice2[index] {
			return false
		}
	}
	return true
}
