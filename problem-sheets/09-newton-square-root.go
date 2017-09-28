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

func NewtonSqrt(number float64, numIterations int) float64 {
	newtonSqrt := 1.0 // guess 1.0 for all numbers
	for i := 0; i < numIterations; i++ {
		lastValue := newtonSqrt
		// compute the next Z once
		newtonSqrt = getNextZ(newtonSqrt, number)
		if newtonSqrt == lastValue {
			break // stop computing when we start getting back the same value.
		}
	}
	return newtonSqrt
}

// function prints a formatted output of a comparison between both Sqrt functions.
func printSqrtComparison(number float64) {
	fmt.Printf("math.Sqrt(%f) = %f\n", number, math.Sqrt(number))
	fmt.Printf("NewtonSqrt(%f) = %f\n", number, NewtonSqrt(number, 10))
	fmt.Println()
}

func main() {
	number := 0.0
	var err error = nil
	for number != -1 { // sentinal value
		fmt.Println("Enter a value to see a sqrt comparisons. Enter -1 to quit.")
		if number, err = util.ReadFloat(); err == nil && number != -1 {
			printSqrtComparison(number)
		}
	}

}

/* Output
Enter a value to see a sqrt comparisons. Enter -1 to quit.
67
math.Sqrt(67.000000) = 8.185353
NewtonSqrt(67.000000) = 8.185353

Enter a value to see a sqrt comparisons. Enter -1 to quit.
4.54
math.Sqrt(4.540000) = 2.130728
NewtonSqrt(4.540000) = 2.130728

Enter a value to see a sqrt comparisons. Enter -1 to quit.
3.14567
math.Sqrt(3.145670) = 1.773604
NewtonSqrt(3.145670) = 1.773604

Enter a value to see a sqrt comparisons. Enter -1 to quit.
543
math.Sqrt(543.000000) = 23.302360
NewtonSqrt(543.000000) = 23.302360

Enter a value to see a sqrt comparisons. Enter -1 to quit.
876
math.Sqrt(876.000000) = 29.597297
NewtonSqrt(876.000000) = 29.597297

Enter a value to see a sqrt comparisons. Enter -1 to quit.
234
math.Sqrt(234.000000) = 15.297059
NewtonSqrt(234.000000) = 15.297059

Enter a value to see a sqrt comparisons. Enter -1 to quit.
-1
*/
