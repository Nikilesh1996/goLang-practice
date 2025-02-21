package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

type Rectange struct {
	width, height float64
}

func (r Rectange) area() float64 {
	return 2.0 * (r.width + r.height)
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return 2.0 * math.Pi * c.area()
}

func exampleOfSwitchCase(s Shape) {
	// Here concreteType is the caster value of the
	switch concreteType := s.(type) {
	case Circle:
		fmt.Printf("This is a circle %v\n", concreteType.radius)
	case Rectange:
		fmt.Printf("This is a rectangle %v %v\n", concreteType.height, concreteType.width)
	default:
		fmt.Println("This is not a valid shape")
	}
}

func main() {
	circle := Circle{
		radius: 20,
	}

	rectangle := Rectange{
		width:  20,
		height: 20,
	}

	exampleOfSwitchCase(circle)
	exampleOfSwitchCase(rectangle)
}
