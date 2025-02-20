package main

import (
	"fmt"
	"math"
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

// Good interface naming
type Copier interface {
	Copy(sourceFile string, destinationFile string) (bytesCopied int)
}

type Shape interface {
	area() float64
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func seeTypeAssertion(shape Shape) {
	// Second param says wether the type is a circle
	circle, isShape := shape.(Circle)
	fmt.Printf("Is given shape a circle %v %v\n", isShape, circle)
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

	// This is type assertion
	circle := Circle{
		radius: 2.0,
	}

	seeTypeAssertion(circle)
}
