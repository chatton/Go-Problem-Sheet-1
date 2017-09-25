/*
@author chatton

Problem Description

Write a function that merges two sorted lists into a new sorted list.
[1,4,6],[2,3,5] → [1,2,3,4,5,6].
*/

package main

import "fmt"

// mergeLists function flattens a list of lists
// into a single list.
func mergeLists(lists ...[]int) []int { // can provide a variable number of arguments using ...
	var result []int
	for _, list := range lists {
		for _, val := range list {
			result = append(result, val)
		}
	}
	return result
}

func main() {
	sortedList1 := []int{1, 2, 3}
	sortedList2 := []int{4, 5, 6}
	sortedList3 := []int{7, 8, 9}
	mergedList := mergeLists(sortedList1, sortedList2)
	fmt.Println("Merging Lists:", sortedList1, "and", sortedList2)
	fmt.Println("Merged List:", mergedList)
	fmt.Println()

	fmt.Println("Merging Lists:", sortedList1, ",", sortedList2, "and", sortedList3)
	fmt.Println("Merged List:", mergeLists(sortedList1, sortedList2, sortedList3))
}

/*
Sample output:

Merging Lists: [1 2 3] and [4 5 6]
Merged List: [1 2 3 4 5 6]

Merging Lists: [1 2 3] , [4 5 6] and [7 8 9]
Merged List: [1 2 3 4 5 6 7 8 9]
*/
