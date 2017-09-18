/*
Problem description.

Write a function that tests whether a string is a palindrome.
*/

package main

import (
    "fmt"
    "strings"
    "./util"
)

func isPalindrome(inputString string) bool {
    asLowerCase := strings.ToLower(inputString)
    return util.Reverse(asLowerCase) == asLowerCase
}

func main(){
    fmt.Println("level is palindrome:", isPalindrome("level"))
    fmt.Println("noon is palindrome:", isPalindrome("noon"))
    fmt.Println("hello is palindrome:", isPalindrome("hello"))
    fmt.Println("world is palindrome:", isPalindrome("world"))
}