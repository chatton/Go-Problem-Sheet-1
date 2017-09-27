/*
@author chatton

Problem description

Write a function that takes a list of integers as input, and returns true if any pair of elements sum to zero, and false otherwise.
*/

package main

import "fmt"

func pairSum(numbers []int) bool {
	for i, num := range numbers {
		for j, innerNum := range numbers {
			if i == j { // it's the same number, don't pair with itself.
				// the only number that will sum to 0 when added to itself is 0
				continue
			}
			if sum := num + innerNum; sum == 0 {
				fmt.Println(num, "and", innerNum, "summed to 0!")
				return true
			}
		}
	}
	return false
}

func main() {
	fmt.Println(pairSum([]int{1, 2, 3, 4, -4}))
	fmt.Println(pairSum([]int{2, -3, 2, 1, 0}))
	fmt.Println(pairSum([]int{1, 1, 0, 0, 1}))
	fmt.Println(pairSum([]int{0, 199, 423, -423, 3}))
}
