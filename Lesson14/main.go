package main

import (
	"errors"
	"fmt"
)

func sendCustomError() error {
	return errors.New("This is a custom error")
}

func main() {
	fmt.Println(sendCustomError().Error())

	// Loops in Go programming language
	for i := 0; i < 10; i++ {
		fmt.Printf("%v ", i)
	}

	fmt.Println()

	// This is a while loop
	index := 0
	for index < 10 {
		fmt.Printf("%v ", index)
		index++
	}

	fmt.Println()
}
