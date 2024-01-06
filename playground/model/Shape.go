package model

// Vertex type
type Vertex struct {
	X, Y float64
}

// pointer receiver: allowing the method to modify the original instance
func (v *Vertex) Scale(f float64) {
	v.X *= f // Go automatically dereferences the pointer for you: (*v).X
	v.Y *= f
}

// Line type
// Defined by its Start and End Vertices.
type Line struct {
	Start, End Vertex
}
