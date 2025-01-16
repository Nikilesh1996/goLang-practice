package main

import "fmt"

// This is a higher order function
func higherOrderFunction(message string, formatter func(string) string) string {
	firstMessage := formatter(message)
	secondMessage := formatter(firstMessage)
	thirdMessage := formatter(secondMessage)
	return thirdMessage
}

func exclamationAdder(message string) string {
	return fmt.Sprintf("%s!", message)
}

func main() {
	message := higherOrderFunction("Hello", exclamationAdder)
	fmt.Println(message)
}
