package main

import (
	"strconv"
	"fmt"
)

func main() {
	// Casting float to integer

	myFloatingType := 2.4
	myIntegerType := int(myFloatingType)

	fmt.Println("Floating variable", myFloatingType)
	fmt.Println("Integer type", myIntegerType)

	fmt.Println("Converting float to string", strconv.FormatFloat(myFloatingType, 'f', -1, 64))
	fmt.Println("Converting int to string", strconv.Itoa(int(myFloatingType)))
}