/*
@author chatton

Problem description

Write a program that prints the numbers from 1 to 100,
except for the following conditions. For multiples of three print "Fizz" instead of the number,
and for the multiples of five print "Buzz".
For numbers which are multiples of both three and five print "FizzBuzz".
*/

package main

import (
	"./util"
	"fmt"
)

// FizzBuzz function prints out every number between the start value (inclusive)
// and the finish value (not-inclusive). If the number is divisible
// by the fizzNum it prints "Fizz", if it's divisible by the
// buzzNum it prints "Buzz", if it's divisible by both it prints
// FizzBuzz. Otherwise, it prints the number itself.
func FizzBuzz(fizzNum, buzzNum, start, finish int) {
	for i := start; i < finish; i++ {
		/*
		   if X % A == 0 and X % B == 0, then X % (A * B) == 0
		   need to check for this condition first otherwise because the
		   "Fizz" and "Buzz" conditions will both evaluate to true.
		*/
		if i%(fizzNum*buzzNum) == 0 {
			fmt.Println("FizzBuzz")
		} else if i%fizzNum == 0 {
			fmt.Println("Fizz")
		} else if i%buzzNum == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}

func main() {

	fmt.Println("Enter Fizz number, Buzz number and finishing point space separated on one line. E.g. [3 5 100]")
	choices, err := util.AsIntSlice(util.ReadLine())
	if err == nil {
		FizzBuzz(choices[0], choices[1], 1, choices[2]+1)
	}
}

/*
Sample output:

Enter Fizz number, Buzz number and finishing point space separated on one line. E.g. [3 5 100]
4 7 120
1
2
3
Fizz
5
6
Buzz
Fizz
9
10
11
Fizz
13
Buzz
15
Fizz
17
18
19
Fizz
Buzz
22
23
Fizz
25
26
27
FizzBuzz
29
30
31
Fizz
33
34
Buzz
Fizz
37
38
39
Fizz
41
Buzz
43
Fizz
45
46
47
Fizz
Buzz
50
51
Fizz
53
54
55
FizzBuzz
57
58
59
Fizz
61
62
Buzz
Fizz
65
66
67
Fizz
69
Buzz
71
Fizz
73
74
75
Fizz
Buzz
78
79
Fizz
81
82
83
FizzBuzz
85
86
87
Fizz
89
90
Buzz
Fizz
93
94
95
Fizz
97
Buzz
99
Fizz
101
102
103
Fizz
Buzz
106
107
Fizz
109
110
111
FizzBuzz
113
114
115
Fizz
117
118
Buzz
Fizz
*/
