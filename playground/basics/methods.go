package basics

import (
	"fmt"
	"playground/model"
)

func MethodBasics() {
	fmt.Println("Method Basics")
	fmt.Println("----------------")

	mutatorMethods()
}

func mutatorMethods() {

	// Create a new rectangle using the constructor
	rect := model.NewRectangle(5.0, 10.0)

	// Scale the rectangle
	// Go allows calling methods on values even when the receiver has a pointer type.
	// implicit referencing: Go will automatically take the address of the value for you.
	rect.Scale(2.0) // pointer receiver, method operates on the rect instance
	(*rect).Scale(3.0)

	// Failed mutator!
	rect.NoScale(1000) // value receiver, method operates on a COPY of rect instance

	// modify and return
	result := rect.ScaleAndReturn(5)
	rect = &result

	// Calculate and print the area
	area := rect.Area()

	fmt.Printf("Rectangle %v area: %v\n", rect, area) // &{150 300} area: 45000

}
