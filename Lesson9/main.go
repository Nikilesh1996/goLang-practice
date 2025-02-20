package main

import "fmt"

type Rectangle struct {
	width  int
	height int
}

// Access structure in a function
// These are methods on a sturct
func (r Rectangle) area() int {
	return r.width * r.height
}

func main() {
	r := Rectangle{
		height: 10,
		width:  10,
	}

	// This is how you call a method in the structure
	fmt.Println(r.area())
}
