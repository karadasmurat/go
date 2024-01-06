package basics

import (
	"fmt"
	"playground/model"
)

func MethodBasics() {
	fmt.Println("Method Basics")
	fmt.Println("----------------")

	vrtx := model.Vertex{X: 10, Y: 20}

	// Go allows calling methods on values even when the receiver has a pointer type.
	// implicit referencing: Go will automatically take the address of the value for you.
	(&vrtx).Scale(2)
	vrtx.Scale(5)

	fmt.Println(vrtx) // {100 200}
}
