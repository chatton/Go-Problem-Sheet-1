package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func CollatzConjecture(num int) {
	fmt.Println("Initial number:", num)
	for {
		lastResult := num

		if num%2 == 1 {
			fmt.Println(num, "Number is odd, multiplying by 3 and adding 1.")
			num *= 3
			num += 1
		} else {
			fmt.Println(num, "Number is even, dividing by 2.")
			num /= 2
		}

		fmt.Println("Resulting number:", num)

		if num == lastResult {
			fmt.Println("Squence started repeating. Stopping.")
			return
		}
	}
}

func numberIsValid(n int) bool {
	return n > 0
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Enter a positive integer.")
		scanner.Scan()
		userInput := scanner.Text()

		if num, err := strconv.Atoi(userInput); err == nil {
			// successful conversion
			if numberIsValid(num) {
				CollatzConjecture(num)
				break
			}
		}
	}

}
