/*
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
	printSqrtComparison(16)
	printSqrtComparison(88)
	printSqrtComparison(100)
	printSqrtComparison(44)
	printSqrtComparison(42.64)
	printSqrtComparison(856.454)
	printSqrtComparison(22.412)
	printSqrtComparison(99.23)
}

/* Output
math.Sqrt(16.000000) = 4.000000
NewtonSqrt(16.000000) = 4.000000

math.Sqrt(88.000000) = 9.380832
NewtonSqrt(88.000000) = 9.380832

math.Sqrt(100.000000) = 10.000000
NewtonSqrt(100.000000) = 10.000000

math.Sqrt(44.000000) = 6.633250
NewtonSqrt(44.000000) = 6.633250

math.Sqrt(42.640000) = 6.529931
NewtonSqrt(42.640000) = 6.529931

math.Sqrt(856.454000) = 29.265235
NewtonSqrt(856.454000) = 29.265235

math.Sqrt(22.412000) = 4.734131
NewtonSqrt(22.412000) = 4.734131

math.Sqrt(99.230000) = 9.961426
NewtonSqrt(99.230000) = 9.961426
*/
