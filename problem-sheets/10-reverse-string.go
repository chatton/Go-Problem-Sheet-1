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
	fmt.Println(util.Reverse("Hello"))
	fmt.Println(util.Reverse("World"))
	fmt.Println(util.Reverse("This string is reversed!"))
	fmt.Println(util.Reverse(util.Reverse("Reversed Twice")))
	fmt.Println(util.Reverse("The quick brown 狐 jumped over the lazy 犬"))
}

/*
Output:

olleH
dlroW
!desrever si gnirts sihT
Reversed Twice
*/
