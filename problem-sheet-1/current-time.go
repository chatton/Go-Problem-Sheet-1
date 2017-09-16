/*
Problem description

Write a program that prints the current time and date to the console.
*/

package main

import (
    "fmt"
    "time"
)

func main(){
    fmt.Println("The current time is," , time.Now());
}