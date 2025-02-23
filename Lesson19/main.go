package main

import (
	"fmt"
)

func sum(elements ...float64) float64 {
	total := 0.0

	for i := 0; i < len(elements); i++ {
		total += elements[i]
	}

	return total
}

func main() {
	total := sum(1.2, 2.2, 3.2)
	fmt.Println(total)
}