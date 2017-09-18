
/*
Write a guessing game where the user must guess a secret number. 
After every guess the program tells the user whether their number was too large or too small. 
At the end the number of tries needed should be printed. 
It counts only as one try if they input the same number multiple times consecutively.
*/

package main

import (
    "fmt"
    "math/rand"
    "time"
)


func displayMessage(secretNum, guessedNum, numGuesses int) {
       if guessedNum == secretNum {
            fmt.Println("You guessed the number correctly! It took you", numGuesses, "guesses")
        } else if guessedNum < secretNum {
            fmt.Println("Go higher!")
        } else {
            fmt.Println("Go lower!")
        }
}

func main() {
    // so we don't get the same number each time
    rand.Seed(time.Now().UTC().UnixNano()) 
    var secretNum = rand.Intn(100) + 1 // don't want to include 0
    var numGuesses int = 0
    var guess int = -1
    for guess != secretNum {
        lastGuess := guess
        
        fmt.Println("Guess a number between 1 and 100")
        fmt.Scan(&guess)
        
        if lastGuess != guess {
            numGuesses++
        }

        displayMessage(secretNum, guess, numGuesses)
    }
}

/* Sample Input/Output

Guess a number between 1 and 100
50
Go lower!
Guess a number between 1 and 100
25
Go higher!
Guess a number between 1 and 100
30
Go lower!
Guess a number between 1 and 100
26
Go higher!
Guess a number between 1 and 100
27
You guessed the number correctly! It took you 5 guesses
*/