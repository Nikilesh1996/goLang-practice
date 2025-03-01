package main

import (
	"fmt"
)

func createMatrix(rows, cols int) [][]int {
	matrix := make([][]int, 0)

	for i := 0; i < rows; i++ {
		row := make([]int, 0)

		for j := 0; j < cols; j++ {
			row = append(row, i*j)
		}

		matrix = append(matrix, row)
	}

	return matrix
}

func indexOfFirstBadWord(message []string, badWords []string) int {
	for i, word := range message {
		for _, badWord := range badWords {
			if word == badWord {
				return i
			}
		}
	}

	return -1
}

func main() {
	matrix := createMatrix(2, 2)

	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("%v ", matrix[i][j])
		}

		fmt.Println()
	}

	fmt.Println()

	messages := []string{"nikilesh", "is", "a", "bad", "boy"}
	badWords := []string{"bad", "idiot"}

	fmt.Println("Index of the first bad word in message: ")
	fmt.Println(indexOfFirstBadWord(messages, badWords))
}
