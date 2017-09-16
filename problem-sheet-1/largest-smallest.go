/*
Problem description

Write a function that returns the largest and smallest elements in a list.
*/

package main

import "fmt"

func largestSmallest(array []int) (int, int){
    var largest int = array[0]
    var smallest int = array[0]

    for i := 0; i < len(array); i++ {
        val := array[i]
        if val > largest {
            largest = val 
        }
        if val < smallest {
            smallest = val
        }
    }

    return largest, smallest

}

func printResults(array []int) {
    largest, smallest := largestSmallest(array)
    fmt.Println("Input array:", array)
    fmt.Println("Largest: ", largest, "Smallest:", smallest)
    fmt.Println()
}

func main() {
    var arr1 = []int {1,2,3,4,5}
    var arr2 = []int {22,-14,-16,42,100}
    var arr3 = []int {-1,543,43,12,53}
    var arr4 = []int {44,33,22,-11,23}
    var arr5 = []int {1,12,23,54,4}
    
    printResults(arr1)
    printResults(arr2)
    printResults(arr3)
    printResults(arr4)
    printResults(arr5)
}