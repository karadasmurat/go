package model

// Point type
type Point struct {
	X, Y float64
}

// Line type
// The Line type has two fields, Begin and End, both of type Point.
type Line struct {
	Begin, End Point
}
