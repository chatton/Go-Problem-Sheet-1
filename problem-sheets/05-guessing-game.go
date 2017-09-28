/*
@author chatton

Write a guessing game where the user must guess a secret number.
After every guess the program tells the user whether their number was too large or too small.
At the end the number of tries needed should be printed.
It counts only as one try if they input the same number multiple times consecutively.
*/

package main

import (
	// fmt is using for Println and also Scan to take in user input.
	"fmt"
	// using math/rand package to generate random numbers.
	"math/rand"
	// used to Seed the rand package.
	//"strconv"
	// to convert from string -> int
	"./util"
	"time"
)

// function to display the correct message to the user
// based on the randomly generated number, the user's guessed number
// and the number of guesses they've made so far.
func displayMessage(secretNum, guessedNum, numGuesses int) { // only need to specify the type of parameters once if they're all the same
	if guessedNum == secretNum {
		fmt.Println("You guessed the number correctly! It took you", numGuesses, "guesses")
	} else if guessedNum < secretNum {
		fmt.Println("Go higher!")
	} else {
		fmt.Println("Go lower!")
	}
}

func getValidGuess() int {
	fmt.Println("Please enter a number between 1 and 100")
	line := util.ReadLine()
	choices, err := util.SplitLine(line)
	for err != nil || len(choices) < 1 || choices[0] < 1 || choices[0] > 100 {
		fmt.Println("Please enter a number between 1 and 100")
		choices, err = util.SplitLine(util.ReadLine())
	}
	return choices[0]

}

func main() {
	// so we don't get the same number each time
	rand.Seed(time.Now().UTC().UnixNano())
	var secretNum = rand.Intn(100) + 1 // don't want to include 0, (full range is 1 - 100)
	var numGuesses int = 0             // the user hasn't guessed anything yet.
	var guess int = -1                 // pick some value that doesn't fall in the possible range.
	for guess != secretNum {           // 'for' can act like a 'while' loop in Go
		lastGuess := guess // to check if the current guess is the same as the last one

		guess = getValidGuess()

		// there is only another guess if it's different from the last one
		if lastGuess != guess {
			numGuesses++
		}

		// print the relevant message to the screen.
		displayMessage(secretNum, guess, numGuesses)
	}
}

/* Sample Input/Output

Please enter a number between 1 and 100
-10
Please enter a number between 1 and 100
101
Please enter a number between 1 and 100
0
Please enter a number between 1 and 100
50
Go higher!
Please enter a number between 1 and 100
75
Go lower!
Please enter a number between 1 and 100
66
Go higher!
Please enter a number between 1 and 100
70
Go higher!
Please enter a number between 1 and 100
73
You guessed the number correctly! It took you 5 guesses
*/
