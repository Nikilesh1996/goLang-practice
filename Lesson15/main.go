package main

import (
	"fmt"
	"strconv"
)

func fizzbuzz() {
	for i := 0; i <= 100; i++ {
		stringToPrint := ""

		if i % 3 == 0 {
			stringToPrint += "fizz"
		}

		if i % 5 == 0 {
			stringToPrint += "buzz"
		}

		if stringToPrint == "" {
			stringToPrint += strconv.Itoa(i)
		}

		fmt.Println(stringToPrint)
	}
}

func main() {
	fizzbuzz()
}