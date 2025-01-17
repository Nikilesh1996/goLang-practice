package main

import "fmt"

// function that returns a function that uses variables outside its scope
func adder() func(int) int {
	sum := 0

	return func(value int) int {
		sum += value
		return sum
	}
}

func demonstrateDefer(value int) string {

	// This gets executed regardless before every return
	defer fmt.Println("This is from the defer statement")

	if value > 20 {
		return "Excellent"
	}

	if value > 10 {
		return "Good"
	}

	return "Ok"
}

func main() {
	harryPotterAdder := adder()
	harryPotterAdder(1)
	harryPotterAdder(2)

	fmt.Println(harryPotterAdder(3))

	// This is for block scope
	sum := 5
	{
		sum := 3
		fmt.Printf("Scoping %d\n", sum)
	}

	fmt.Printf("Scoping %d\n", sum)

	fmt.Println(demonstrateDefer(10))
	fmt.Println(demonstrateDefer(20))
	fmt.Println(demonstrateDefer(22))
}
