package util

func Reverse(stringToReverse string) string {
    // based on this SO answer
    // https://stackoverflow.com/questions/1752414/how-to-reverse-a-string-in-go
    runes := []rune(stringToReverse) // make an array full of the characters in the string
    for i, rune := range stringToReverse {
        // starts out at the length of the array
        // -1 to start out at the last position.
        // -i to move backwards through the positions.
        runes[len(runes) - 1 - i] = rune
    }
    return string(runes) // convert the rune array back into a string.
}