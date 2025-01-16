package main

import "fmt"

func thisIsAFunctionWithATupleReturnType(x int, y int) (int, int) {
	return x + y, x - y
}

func thisIsAFunctionWithOneReturnType(x int, y int) int {
	return x + y
}

// This is a function with no return types
func thisIsAFunctionWithNoReturnTypes(x int, y int) {
	fmt.Printf("%v is the value of x, and %v is the value of y\n", x, y)
}

func main() {
	const thisIsAConstant = "constant"
	statement := fmt.Sprintf("This is a statement with a constant: %v\n", thisIsAConstant)

	fmt.Printf(statement)
	thisIsAFunctionWithNoReturnTypes(1, 2)
	sum := thisIsAFunctionWithOneReturnType(1, 2)

	fmt.Printf("This is the sum of the function: %v\n", sum)

	total, deducted := thisIsAFunctionWithATupleReturnType(1, 2)
	fmt.Printf("This is the sum %v and this is the difference %v\n", total, deducted)

	// This can ignore the value
	_ = thisIsAFunctionWithOneReturnType(2, 3)

	// _ := "Nikilesh" is not possible
}
