package basics

import (
	"fmt"
	"math"
)

// Mover is an interface representing animals that can move
type Mover interface {
	Move()
}

// Runner is a concrete implementation of Mover for animals that run
type Runner struct{}

// Move implements the Move method for Runner
func (r Runner) Move() {
	fmt.Println("Running")
}

// Crawler is a concrete implementation of Mover for animals that crawl
type Crawler struct{}

// Move implements the Move method for Crawler
func (c Crawler) Move() {
	fmt.Println("Crawling")
}

// Animal is a struct representing an animal with a name and a Mover
type Animal struct {
	Name  string
	Mover // Embedded Mover interface
}

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

	// An instance of the concrete implementation of Mover Interface.
	runner := Runner{}

	// Create a Cheetah using the Runner as its mover
	cheetah := Animal{Name: "Cheetos", Mover: runner}

	// Create a Lion using the same Runner as its mover
	lion := Animal{Name: "Marty", Mover: runner}

	// A slice of Animal
	animals := []Animal{cheetah, lion}
	for _, v := range animals {
		v.Move()
	}

	// A slice of Mover!
	// when we embed an interface type, its fields and methods are directly accessible to the outer struct.
	// therefore, we can use it anywhere the Interface type is expected.
	movers := []Mover{cheetah, lion}
	for _, v := range movers {
		v.Move()
	}

}
