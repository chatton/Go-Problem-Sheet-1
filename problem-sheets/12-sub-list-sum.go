/*
@author: chatton

Problem description

Write a function that takes a list of integers as input, and returns true if any number of elements sum to zero, and false otherwise.
*/

package main

import "fmt"

// simple function to sum up the values in an integer list
func sum(nums []int) int {
	tally := 0
	for _, num := range nums {
		tally += num
	}
	return tally
}

func subListSum(nums []int, target int, partial []int) bool {
	sumOfNumbers := sum(partial)
	// I consulted this SO question https://stackoverflow.com/questions/4632322/finding-all-possible-combinations-of-numbers-to-reach-a-given-sum
	// to find out how to get every possible permutation of values.

	if sumOfNumbers == target && len(partial) > 0 { // added length condition so we don't sum to 0 by default when we sum an empty list.
		fmt.Println("Summed to", target, "using combination", partial)
		return true // it is possible to sum to the target with the given values.
	}

	for i := 0; i < len(nums); i++ {
		remaining := nums[i+1:] // getting every element after this one
		nextPartial := append(partial, nums[i])
		// don't want to return the result directly, want to evaluate everything until we get at true return value.
		if reachedTarget := subListSum(remaining, target, nextPartial); reachedTarget {
			return reachedTarget // stop as soon as we find out it is possible, no need to evaluate any more combinations.
		}

	}

	// every combination has been attempted, it is not possible to sum to the target with the given list.
	return false
}

func main() {
	fmt.Println(subListSum([]int{1, 2, 2, -3}, 0, []int{}))
	fmt.Println(subListSum([]int{1, 2, 2, 4}, 0, []int{}))
	fmt.Println(subListSum([]int{-100, 50, 75, 25, -51}, 0, []int{}))
	fmt.Println(subListSum([]int{-100, 50, 75, 25, -50}, 0, []int{}))
}

/*
sample output:

Summed to 0 using combination [1 2 -3]
true
false
Summed to 0 using combination [-100 75 25]
true
Summed to 0 using combination [-100 50 75 25 -50]
true
*/
