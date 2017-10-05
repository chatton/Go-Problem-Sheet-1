/*
@author chatton

Problem Description

Write a function that merges two sorted lists into a new sorted list.
[1,4,6],[2,3,5] → [1,2,3,4,5,6].
*/

package main

import (
	"./util"
	"fmt"
	"sort"    // package for sorting values
	"strings" // for string conversions.
)

// mergeLists function flattens a list of lists
// into a single sorted list.
func mergeLists(lists ...[]float64) []float64 { // can provide a variable number of arguments using ...
	var result []float64
	// go through every list
	for _, list := range lists {
		// append takes a variable numbers of arguments
		// using list... passes in the contents of the list as these arguments
		result = append(result, list...)
	}
	// sorts the slice of float64s
	sort.Float64s(result)
	return result
}

func main() {
	// allow user to enter floats.
	listOfLists := [][]float64{}
	numLists, err := util.ReadInt("How many lists do you want to merge?")
	for err != nil {
		numLists, err = util.ReadInt("Enter a valid number.")
	}
	for len(listOfLists) < numLists {
		// allow for any number of spaces between inputs, e.g. "12.0  43.23    -12.42   32.23  "
		line := util.ReadLine("Enter space separated numbers")
		line = strings.Join(strings.Fields(line), " ")        // rejoin the list into a single string
		if list, err := util.AsFloatSlice(line); err == nil { // get the actual integer values from the string.
			fmt.Println("Adding list", list)
			listOfLists = append(listOfLists, list)
		} else {
			fmt.Println("Numbers entered were invalid. Make sure to enter a space separated list of integers.")
		}
	}
	fmt.Println("Merged and sorted list: ", mergeLists(listOfLists...))

}

/*
Sample output/input:

ow many lists do you want to merge?

nter space separated numbers
2.43 12.43 123.32
dding list [32.43 12.43 123.32]
nter space separated numbers
123.4 12 43.2
dding list [-123.4 12 43.2]
nter space separated numbers
2.2 23.23 0.232
dding list [32.2 23.23 0.232]
nter space separated numbers
2
dding list [32]
erged and sorted list:  [-123.4 0.232 12 12.43 23.23 32 32.2 32.43 43.2 123.32]
*/
