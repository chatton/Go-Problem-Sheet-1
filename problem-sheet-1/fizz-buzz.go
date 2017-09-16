/*
Problem description

Write a program that prints the numbers from 1 to 100, 
except for the following conditions. For multiples of three print "Fizz" instead of the number,
and for the multiples of five print "Buzz".
 For numbers which are multiples of both three and five print "FizzBuzz".
*/

 package main

 import "fmt"

func FizzBuzz(fizzNum int, buzzNum int, start int, finish int){
    for i := start; i < finish; i++ {
        if i % (fizzNum * buzzNum) == 0 {
            fmt.Println("FizzBuzz")
        } else if i % fizzNum == 0 {
            fmt.Println("Fizz")
        } else if i % buzzNum == 0 {
            fmt.Println("Buzz")
        } else {
            fmt.Println(i)
        }
    }
}

 func main(){
    FizzBuzz(3, 5, 1, 101);
 }