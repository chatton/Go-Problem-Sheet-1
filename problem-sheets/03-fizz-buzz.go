/*
Problem description

Write a program that prints the numbers from 1 to 100,
except for the following conditions. For multiples of three print "Fizz" instead of the number,
and for the multiples of five print "Buzz".
For numbers which are multiples of both three and five print "FizzBuzz".
*/

package main

import "fmt"

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
	FizzBuzz(3, 5, 1, 101)
}

/*
Sample output:

1
2
Fizz
4
Buzz
Fizz
7
8
Fizz
Buzz
11
Fizz
13
14
FizzBuzz
16
17
Fizz
19
Buzz
Fizz
22
23
Fizz
Buzz
26
Fizz
28
29
FizzBuzz
31
32
Fizz
34
Buzz
Fizz
37
38
Fizz
Buzz
41
Fizz
43
44
FizzBuzz
46
47
Fizz
49
Buzz
Fizz
52
53
Fizz
Buzz
56
Fizz
58
59
FizzBuzz
61
62
Fizz
64
Buzz
Fizz
67
68
Fizz
Buzz
71
Fizz
73
74
FizzBuzz
76
77
Fizz
79
Buzz
Fizz
82
83
Fizz
Buzz
86
Fizz
88
89
FizzBuzz
91
92
Fizz
94
Buzz
Fizz
97
98
Fizz
Buzz
*/
