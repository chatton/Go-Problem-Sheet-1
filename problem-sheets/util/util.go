/*
@author chatton
*/

package util

import (
	"bufio"
	"os"
	"strconv"
    "strings"
    "unicode/utf8"

)

// can re-use in the different parsing functions instead of creating new one every time
var scanner *bufio.Scanner

func Reverse(stringToReverse string) string {
	// I consulted this SO question
	// https://stackoverflow.com/questions/1752414/how-to-reverse-a-string-in-go
	// in order to figure out how to account for the encodings in the reversal.

    // get the number of runes in the string, default len function
    // gives back the number of bytes, depending on the encoding a rune could be
    // a different number of bytes.
    numRunes := utf8.RuneCountInString(stringToReverse)
    runes := []rune(stringToReverse)
    // reverse the values in the rune slice.
    for i := 0; i < num/2; i++ {
        runes[i], runes[numRunes-1-i] = runes[numRunes-1-i], runes[i]
    }
    return string(runes) // convert the slice back into a string.
}

// simple function to sum up the values in an integer list
func Sum(nums []int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

// simple function to read a line in from the console.
func ReadLine() string {
    if scanner == nil { 
        scanner = bufio.NewScanner(os.Stdin)
    }
    
    scanner.Scan()
    return scanner.Text()
    
}

func ReadInt() (int, error) {
    return strconv.Atoi(ReadLine())
}

func ReadFloat() (float64, error) {
    return strconv.ParseFloat(ReadLine(), 64)
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
