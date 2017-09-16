package util

func Reverse(stringToReverse string) string {
    // based on this SO answer
    // https://stackoverflow.com/questions/1752414/how-to-reverse-a-string-in-go
    length := len(stringToReverse)
    runes := make([]rune, length) // make an array big enough to store all the characters
    index := length
    for _, rune := range stringToReverse {
        index--
        runes[index] = rune
    }
    return string(runes)
}