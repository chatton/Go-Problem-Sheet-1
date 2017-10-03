/*
@author chatton

Problem description

Write a function to reverse a string in Go.
*/

package main

import (
	// for the Reverse function.
	"./util"
	"fmt"
)

func main() {

	word := util.ReadLine("Enter word to reverse. (q to quit)")
	for word != "q" {
		fmt.Println("Reversed:", util.Reverse(word))

		fmt.Println()
		word = util.ReadLine("Enter word to reverse. (q to quit)")
	}
}

/*
Output:

Enter word to reverse. (q to quit)
hello
Reversed: olleh
Enter word to reverse. (q to quit)
hello world
Reversed: dlrow olleh
Enter word to reverse. (q to quit)
what time is it?
Reversed: ?ti si emit tahw
Enter word to reverse. (q to quit)
q
*/
