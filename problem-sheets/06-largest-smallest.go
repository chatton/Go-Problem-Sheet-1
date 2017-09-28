/*
@author chatton

Problem description

Write a function that returns the largest and smallest elements in a list.
*/

package main

import (
	"./util"
	"fmt"
)

// takes in an int slice and returns 2 integers, the largest and smallest from the slice.
func largestSmallest(array []int) (int, int) {
	if len(array) == 0 {
		panic("Invalid argument provided: Array must have length > 0")
	}
	// assume the first element is both the largest and the smallest.
	var largest int = array[0]
	var smallest int = array[0]

	// look at every element
	for i := 0; i < len(array); i++ {
		val := array[i] // get the current value
		// check if it's bigger than the current largest and make it the new largest
		if val > largest {
			largest = val
		}
		// check if it's smaller than current smallest and make it the new smallest
		if val < smallest {
			smallest = val
		}
	}

	// in Go we can have multiple return values.
	return largest, smallest

}

// function to display the results of a call to largestSmallest nicely.
func printResults(array []int) {
	largest, smallest := largestSmallest(array)
	fmt.Println("Input array:", array)
	fmt.Println("Largest: ", largest, "Smallest:", smallest)
	fmt.Println()
}

func main() {
	fmt.Println("Please enter 1 or more space separated integers on a single line.")
	line := util.ReadLine()
	if arr, err := util.SplitLine(line); err == nil {
		printResults(arr)
	}
}

/*
Sample input/output:

Please enter 1 or more space separated numbers on a single line.
1 2 3 4 5 6 7 -123 -42 23
Input array: [1 2 3 4 5 6 7 -123 -42 23]
Largest:  23 Smallest: -123

*/
