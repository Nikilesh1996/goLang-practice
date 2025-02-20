package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
	perimeter() float64
}

// Rectangle implements both the methods in the interface
// This is considered as a type of Shape.
type Rectangle struct {
	width, height float64
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func (r Rectangle) perimeter() float64 {
	return 2 * (r.width + r.height)
}

// Same thing with the circle
type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}


// These methods accept a Shape type 
func calculateArea(shape Shape) float64 {
	return shape.area()
}

func calculatePerimeter(shape Shape) float64 {
	return shape.perimeter()
}

func main() {
	circle := Circle{
		radius: 2,
	}

	rectangle := Rectangle{
		width:  2,
		height: 2,
	}

	fmt.Println(calculateArea(circle))
	fmt.Println(calculateArea(rectangle))

	fmt.Println(calculatePerimeter(circle))
	fmt.Println(calculatePerimeter(rectangle))
}
