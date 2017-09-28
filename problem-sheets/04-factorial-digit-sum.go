/*
@author chatton

Problem description

Write a function that calculates the sum of the digits of the factorial of a number.
n! means n x (n − 1) ... x 3 x 2 x 1. For example, 10! = 10 x 9 x ... x 3 x 2 x 1 = 3628800,
and the sum of the digits in the number 10! is 3 + 6 + 2 + 8 + 8 + 0 + 0 = 27.
Find the sum of the digits in the number 100!
*/

package main

import (
	"fmt"
	// "math/big" is needed to store the result of 100!
	"math/big"
	// "strings" is needed for "Split"
	"strings"
	// "strconv" is needed to parse a string into an int using "Atoi"
	"strconv"
	// for ReadIn and AsIntSlice
	"./util"
)

// function takes in a string, and returns the sum of the digits.
func sumDigitsInString(resultString string) int {
	// split on an empty string to convert a string into a list of characters
	// strings.Split("abcd", ""); returns ["a", "b", "c", "d"]
	allDigits := strings.Split(resultString, "") // get a list of every digit

	totalSum := 0
	for _, digit := range allDigits { // we don't care about the index, only the digit itself.
		val, _ := strconv.Atoi(digit) // parse digit to value "9" -> 9
		totalSum += val               // tally up the total value
	}
	return totalSum
}

func main() {
	fmt.Println("Enter a number to get the sum of the digits in its factorial.")
	num, err := util.ReadInt()
	for err != nil {
		fmt.Println("Invalid number, enter again.")
		num, err = util.ReadInt()
	}

	// we need to use bit.Int because we can't hold the value with standard types
	bigIntptr := new(big.Int)
	// don't need to write a factorial method, we can use MulRange to get the value.
	factorialResult := bigIntptr.MulRange(1, int64(num)) // equivalent to N x (N - 1) x (N - 2) x ... x 1
	totalSum := sumDigitsInString(factorialResult.String())
	fmt.Printf("The sum of the digits of %d! is: %d", num, totalSum)
}

/*
Sample output:

Enter a number to get the sum of the digits in its factorial.
gfd
Invalid number, enter again.
werw
Invalid number, enter again.
125
The sum of the digits of 125! is: 855
*/
