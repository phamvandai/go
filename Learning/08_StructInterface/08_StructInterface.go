package main

import (
	"fmt"
)

// Rectangle struct
type Rectangle struct {
	width float64
	heigh float64
}

// Area of Rectangle
func Area(r *Rectangle) float64 {
	return r.heigh * r.width
}

// Method of struct
func (r *Rectangle) area() float64 {
	return r.heigh * r.width
}

// Circle struct
type Circle struct {
	radius float64
}

// Circle area method
func (c *Circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

// Shape interface
type Shape interface {
	area() float64
}

// Function using interface, can pass arguments of different types, but having
// same interface
func totalArea(shape ...Shape) float64 {
	total := 0.0
	for _, s := range shape {
		total += s.area()
	}
	return total
}

func main() {
	rect1 := Rectangle{width: 2.5, heigh: 3.3}

	fmt.Println("rect1 area", Area(&rect1))
	fmt.Println("rect1 area method", rect1.area())

	rect2 := Rectangle{width: 3.3, heigh: 1.5}
	cir1 := Circle{radius: 2}

	fmt.Println("Total area", totalArea(&rect2, &cir1))
}
