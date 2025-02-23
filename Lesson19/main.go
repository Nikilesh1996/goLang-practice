package main

import (
	"fmt"
)

type Cost struct {
	day int 
	value float64
}

func getCostsByDay(costs []Cost) []float64 {
	costsByDay := []float64{}

	for i := 0; i < len(costs); i++ {
		cost := costs[i]

		for cost.day >= len(costsByDay) {
			costsByDay = append(costsByDay, 0.0)			
		}

		costsByDay[cost.day] += cost.value
	}

	return costsByDay
}

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

	costs := getCostsByDay([]Cost {
		{1, 2.2},
		{2, 2.2},
		{1, 1.1},
	})

	fmt.Println(costs)
}