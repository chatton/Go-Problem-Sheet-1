/*
@author chatton

Problem Description

Write a function that merges two sorted lists into a new sorted list.
[1,4,6],[2,3,5] → [1,2,3,4,5,6].
*/

package main

import (
	"fmt"
	// package for sorting values
	"sort"
)

// mergeLists function flattens a list of lists
// into a single sorted list.
func mergeLists(lists ...[]int) []int { // can provide a variable number of arguments using ...
	var result []int
	// go through every list
	for _, list := range lists {
		// append takes a variable numbers of arguments
		// using list... passes in the contents of the list as these arguments
		result = append(result, list...)
	}
	// sorts the list of integers
	sort.Ints(result)
	return result
}

func main() {
	sortedList1 := []int{1, 4, 6}
	sortedList2 := []int{2, 3, 5}
	sortedList3 := []int{-1, 2, 22}
	mergedList := mergeLists(sortedList1, sortedList2)
	fmt.Println("Merging Lists:", sortedList1, "and", sortedList2)
	fmt.Println("Merged List:", mergedList)
	fmt.Println()

	fmt.Println("Merging Lists:", sortedList1, ",", sortedList2, "and", sortedList3)
	fmt.Println("Merged List:", mergeLists(sortedList1, sortedList2, sortedList3))
}

/*
Sample output:

Merging Lists: [1 4 6] and [2 3 5]
Merged List: [1 2 3 4 5 6]

Merging Lists: [1 4 6] , [2 3 5] and [-1 2 22]
Merged List: [-1 1 2 2 3 4 5 6 22]

*/
