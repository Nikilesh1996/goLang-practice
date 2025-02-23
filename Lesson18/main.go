package main

import (
	"fmt"
)

func main() {
	myArray := [3]int {1, 2, 3}
	var myFloats [len(myArray)]float64

	fmt.Println(len(myFloats))
	dynamicArray := make([]int, 5)

	for i := 0; i < 20; i++ {
		if i < len(dynamicArray) {
			dynamicArray[i] = i
		} else {
			dynamicArray = append(dynamicArray, i)
		}
	}

	for i := 0; i < len(dynamicArray); i++ {
		fmt.Printf("%v ", dynamicArray[i])
	}

	fmt.Println()
	
	jack := []int{}
	for i := 0; i < 20; i++ {
		jack = append(jack, i)
	}

	for i := 0; i < len(jack); i++ {
		fmt.Printf("%v ", i)
	}

	fmt.Println()
}