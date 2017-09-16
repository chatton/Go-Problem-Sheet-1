/*
Problem Description

Write a function that merges two sorted lists into a new sorted list. 
[1,4,6],[2,3,5] → [1,2,3,4,5,6].
*/

package main

import "fmt"


func mergeLists(list1, list2 []int) []int {
    var result []int

    for _, element := range list1 {
        result = append(result, element)
    }

    for _, element := range list2 {
        result = append(result, element)
    }
    return result
}

func main(){
    sortedList1 := []int{1,2,3}
    sortedList2 := []int{4,5,6}
    mergedList := mergeLists(sortedList1, sortedList2)
    fmt.Println("Merging Lists:", sortedList1, "and", sortedList2)
    fmt.Println("Merged List:", mergedList)

}