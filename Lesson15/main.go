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

func printPrimes(max int) {
	for n := 2; n < max + 1; n++ {
		if n == 2 {
			fmt.Println(n)
			continue
		}

		if n % 2 != 0 {
			isPrime := true
			for i := 3; i * i < n + 1; i++ {
				if n % i == 0 {
					isPrime = false
					break
				}
			}

			if !isPrime {
				continue
			}

			fmt.Println(n)
		}
	}
}

func main() {
	fizzbuzz()
	printPrimes(10)
}