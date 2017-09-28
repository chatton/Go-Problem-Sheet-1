/*
@author chatton
*/

package util

import (
	"bufio"
	"os"
	"strconv"
    "strings"
)

func Reverse(stringToReverse string) string {
	// I consulted this SO question
	// https://stackoverflow.com/questions/1752414/how-to-reverse-a-string-in-go
	// in order to figure out how to account for the encodings in the reversal.

	// the reason you can't use len to get the number of characters
	// is because it gives back the number of bytes, and different
	// characters can give back different numbers of bytes.

	num := 0
	runes := []rune(stringToReverse)    // creates a new rune list that from the input string
	for _, r := range stringToReverse { // don't care about the index, just the rune itself
		runes[num] = r
		num++ // the actual number of characters
	}

	// Reverse the contents of the array by
	// iterating to the halfway point and swapping elements
	// on the furthest ends as you go.
	for i := 0; i < num/2; i++ {
		runes[i], runes[num-1-i] = runes[num-1-i], runes[i]
	}

	// Finally, convert back to UTF-8.
	return string(runes)
}

// simple function to sum up the values in an integer list
func Sum(nums []int) int {
	tally := 0
	for _, num := range nums {
		tally += num
	}
	return tally
}

// simple function to read a line in from the console.
func ReadLine() string {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    return scanner.Text()
    
}

func ReadInt() (int, error) {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    return strconv.Atoi(scanner.Text())
}

// functions takes a space seprated string and splits it up into
//a slice of integers 
func AsIntSlice(line string) ([]int, error) {
    splitLine := strings.Split(line, " ")
    result := []int{}
    for _, char := range splitLine {
        if num, err := strconv.Atoi(char); err == nil {
            result = append(result, num)
        } else{ // there was an error
            return []int{}, err
        }
    }
    return result, nil // no error, successful
}
