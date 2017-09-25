/*
Problem description

Write a function that returns the largest and smallest elements in a list.
*/

package main

import "fmt"

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
	var arr1 = []int{1, 2, 3, 4, 5}
	var arr2 = []int{22, -14, -16, 42, 100}
	var arr3 = []int{-1, 543, 43, 12, 53}
	var arr4 = []int{44, 33, 22, -11, 23}
	var arr5 = []int{1, 12, 23, 54, 4}

	printResults(arr1)
	printResults(arr2)
	printResults(arr3)
	printResults(arr4)
	printResults(arr5)
}

/*
Sample output:

Input array: [1 2 3 4 5]
Largest:  5 Smallest: 1

Input array: [22 -14 -16 42 100]
Largest:  100 Smallest: -16

Input array: [-1 543 43 12 53]
Largest:  543 Smallest: -1

Input array: [44 33 22 -11 23]
Largest:  44 Smallest: -11

Input array: [1 12 23 54 4]
Largest:  54 Smallest: 1

*/
