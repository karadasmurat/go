package basics

import (
	"fmt"
	"math"
)

// Shape is an interface for geometric shapes
type Shape interface {
	Area() float64
}

// Circle is a type representing a circle
type Circle struct {
	Radius float64
}

// Implementing the Area method for Circle
func (c Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}

// Triangle is a type representing a triangle
type Triangle struct {
	Base, Height float64
}

// Implementing the Area method for Triangle
func (t Triangle) Area() float64 {
	return t.Base * t.Height / 2
}

// A function, accepts any Shape to print its area
func printArea(s Shape) {
	fmt.Printf("The area of provided shape is: %.2f\n", s.Area())
}

func InterfaceBasics() {
	fmt.Println("Interface Basics")
	fmt.Println("----------------")

	myCircle := Circle{Radius: 10}
	myTriangle := Triangle{Base: 10, Height: 20}

	// Once a type implements an interface, values of that type can be used
	// wherever the interface type is expected.
	printArea(myTriangle) // Polymorphic behavior based on the Shape interface
	printArea(myCircle)   // Polymorphic behavior based on the Shape interface
}
