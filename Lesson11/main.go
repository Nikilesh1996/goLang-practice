package main

import (
	"fmt"
)

type Expense interface {
	cost() float64
}

type Printer interface {
	print()
}

type Email struct {
	isSubscibed bool
	body        string
}

func (e Email) cost() float64 {
	if !e.isSubscibed {
		return 0.05 * float64(len(e.body))
	}

	return 0.01 * float64(len(e.body))
}

func (e Email) print() {
	fmt.Println(e.body)
}

func print(expense Expense, printer Printer) {
	printer.print()
	fmt.Println(expense.cost())
}

func main() {
	email1 := Email{
		isSubscibed: false,
		body:        "This is the first body",
	}

	email2 := Email{
		isSubscibed: true,
		body:        "This is the second body",
	}

	print(email1, email1)
	print(email2, email2)
}
