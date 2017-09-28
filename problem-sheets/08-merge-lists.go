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
	"./util"
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
	fmt.Println("How many lists do you want to merge?")
	listOfLists := [][]int{}
	numLists, err := util.ReadInt()
	for err != nil {
		fmt.Println("Enter a valid number.")
		numLists, err = util.ReadInt()
	}
	for len(listOfLists) < numLists {
		fmt.Println("Enter space separated integers")
		if list, err := util.AsIntSlice(util.ReadLine()); err == nil {
			fmt.Println("Adding list", list)
			listOfLists = append(listOfLists, list)
		}
	}
	fmt.Println("Merged and sorted list: ", mergeLists(listOfLists...))

}

/*
Sample output/input:

How many lists do you want to merge?
3
Enter space separated integers
1 4 9
Adding list [1 4 9]
Enter space separated integers
3 6 0 23 4
Adding list [3 6 0 23 4]
Enter space separated integers
5 9 2 3 4
Adding list [5 9 2 3 4]
Merged and sorted list:  [0 1 2 3 3 4 4 4 5 6 9 9 23]

*/
