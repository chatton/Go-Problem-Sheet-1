/*
@author chatton

Problem description

Write a program that prints the current time and date to the console.
*/

package main

import (
	"fmt"
	// import the 'time' package for the time.Now function which gives back the current time
	"time"
)

func main() {
	fmt.Println("The current time is,", time.Now())
}

/*
Sample output:

The current time is, 2017-09-18 10:12:10.4420887 +0100 BST
*/
