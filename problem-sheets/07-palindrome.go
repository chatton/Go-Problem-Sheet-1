/*
@author chatton

Problem description.

Write a function that tests whether a string is a palindrome.
*/

package main

import (
	"fmt"
	// using strings for the ToLower function
	"strings"
	// for the Reverse function
	"./util"
)

func isPalindrome(inputString string) bool {
	// function checks case insensitively
	asLowerCase := strings.ToLower(inputString)
	return util.Reverse(asLowerCase) == asLowerCase
}

func main() {
	line := util.ReadLine("Enter a string to be tested as a palindrome.")
	fmt.Println(line, "is palindrome:", isPalindrome(line))
}

/*
Output:

level is palindrome: true
noon is palindrome: true
hello is palindrome: false
world is palindrome: false
*/
