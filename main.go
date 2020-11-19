package main

import "fmt"

var currentNumber int

func main() {
	currentNumber = 1
	text := "Hello "

	printWithNumber(text)
	printWithNumber(text)
	printWithNumber(text)
	printWithNumber(text)
}

func printWithNumber(text string) {
	fmt.Println(text, currentNumber)
	fmt.Println("")
	currentNumber = currentNumber + 1
}
