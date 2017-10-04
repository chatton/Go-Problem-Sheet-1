/*
@author chatton

Problem description

Implement the square root function using Newton's method.
In this case, Newton's method is to approximate sqrt(x) by picking a starting point z and then repeating:

z_next = z - ((z*z - x) / (2 * z))
To begin with, just repeat that calculation 10 times and see how close you get to the answer for various values
(1, 2, 3, ...). Next, change the loop condition to stop once the value has stopped changing
(or only changes by a very small delta). How close are you to the math.Sqrt value?
*/

package main

import (
	"./util"
	"fmt"
	"math"
)

func getNextZ(z, x float64) float64 {
	return z - ((z*z - x) / (2 * z))
}

func NewtonSqrt(number float64, maxIterations int) (float64, int) {
	newtonSqrt := 1.0  // guess 1.0 for all numbers
	numIterations := 0 // keep track of how many times we iterated through
	for i := 0; i < maxIterations; i++ {
		lastValue := newtonSqrt
		// compute the next Z once
		newtonSqrt = getNextZ(newtonSqrt, number)
		if newtonSqrt == lastValue {
			break // stop computing when we start getting back the same value.
		}
		numIterations++
	}
	return newtonSqrt, numIterations
}

// function prints a formatted output of a comparison between both Sqrt functions.
func printSqrtComparison(number float64) {
	fmt.Printf("math.Sqrt(%f) = %f\n", number, math.Sqrt(number))
	newtonNum, numIterations := NewtonSqrt(number, 100)
	fmt.Printf("NewtonSqrt(%f) = %f in %d iterations.\n", number, newtonNum, numIterations)
	fmt.Println()
}

func main() {
	number := 0.0
	var err error = nil
	for number != -1 { // sentinal value
		if number, err = util.ReadFloat("Enter a value to see a sqrt comparisons. Enter -1 to quit."); err == nil && number != -1 {
			printSqrtComparison(number)
		}
	}

}

/* Input / Output
Enter a value to see a sqrt comparisons. Enter -1 to quit.
67
math.Sqrt(67.000000) = 8.185353
NewtonSqrt(67.000000) = 8.185353 in 8 iterations.

Enter a value to see a sqrt comparisons. Enter -1 to quit.
4.54
math.Sqrt(4.540000) = 2.130728
NewtonSqrt(4.540000) = 2.130728 in 6 iterations.

Enter a value to see a sqrt comparisons. Enter -1 to quit.
3.14567
math.Sqrt(3.145670) = 1.773604
NewtonSqrt(3.145670) = 1.773604 in 100 iterations.

Enter a value to see a sqrt comparisons. Enter -1 to quit.
542
math.Sqrt(542.000000) = 23.280893
NewtonSqrt(542.000000) = 23.280893 in 9 iterations.

Enter a value to see a sqrt comparisons. Enter -1 to quit.
876
math.Sqrt(876.000000) = 29.597297
NewtonSqrt(876.000000) = 29.597297 in 10 iterations.

Enter a value to see a sqrt comparisons. Enter -1 to quit.
gdfgdf
Enter a value to see a sqrt comparisons. Enter -1 to quit.
666
math.Sqrt(666.000000) = 25.806976
NewtonSqrt(666.000000) = 25.806976 in 9 iterations.

Enter a value to see a sqrt comparisons. Enter -1 to quit.
-1
*/
