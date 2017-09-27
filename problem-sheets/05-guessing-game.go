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
	"bufio"
	// for budio.Scanner to read user input
	"os"
	// to get os.Stdin to read from keyboard
	"strconv"
	// to convert from string -> int
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

func getValidGuess(scanner *bufio.Scanner) int {
	value := -1
	for value <= 0 || value > 100 {
		fmt.Println("Guess a number between 1 and 100")
		scanner.Scan()               // read the next line
		userChoice := scanner.Text() // get the text from the line

		// try and parse an int from the string, if there's an error, display message and prompt again.
		if val, err := strconv.Atoi(userChoice); err == nil {
			if val < 1 || val > 100 { // a valid number was entered, check if it's within the correct range for the game.
				fmt.Println("Please enter a number between 1 and 100")
				continue
			}
			return val // a valid guess
		}
		// the number wasn't valid, there was an error converting from string to int.
		fmt.Println("Please enter a valid number.")
	}
	return -1 // should never get here
}

func main() {
	// so we don't get the same number each time
	scanner := bufio.NewScanner(os.Stdin) // use to take in user input, only make one and re-use
	rand.Seed(time.Now().UTC().UnixNano())
	var secretNum = rand.Intn(100) + 1 // don't want to include 0, (full range is 1 - 100)
	var numGuesses int = 0             // the user hasn't guessed anything yet.
	var guess int = -1                 // pick some value that doesn't fall in the possible range.
	for guess != secretNum {           // 'for' can act like a 'while' loop in Go
		lastGuess := guess // to check if the current guess is the same as the last one

		guess = getValidGuess(scanner)

		// there is only another guess if it's different from the last one
		if lastGuess != guess {
			numGuesses++
		}

		// print the relevant message to the screen.
		displayMessage(secretNum, guess, numGuesses)
	}
}

/* Sample Input/Output

Guess a number between 1 and 100
a
Please enter a valid number.
Guess a number between 1 and 100
101
Please enter a number between 1 and 100
Guess a number between 1 and 100
50
Go lower!
Guess a number between 1 and 100
25
Go lower!
Guess a number between 1 and 100
kjhgjhgj
Please enter a valid number.
Guess a number between 1 and 100
15
Go lower!
Guess a number between 1 and 100
10
Go lower!
Guess a number between 1 and 100
5
Go lower!
Guess a number between 1 and 100
4
You guessed the number correctly! It took you 6 guesses
*/
