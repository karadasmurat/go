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

// Rectangle represents a geometric rectangle
type Rectangle struct {
	Width  float64
	Height float64
}

// Mutator method (modifies the state of the struct)
// Method receives a reference to the original instance this method is called on.
// Scale multiplies the width and height of the rectangle by a factor
func (r *Rectangle) Scale(factor float64) {
	r.Width = r.Width * factor
	r.Height = r.Height * factor
}

// Failed Mutator!
// With a value receiver, the NoScale method operates on a COPY of the original Rectangle value.
func (r Rectangle) NoScale(factor float64) {
	r.Width = r.Width * factor
	r.Height = r.Height * factor
}

// Return new using the value receiver (copy of the instance this method is called on)
// which seems to modify but actually create new with the desired changes!
func (r Rectangle) ScaleAndReturn(factor float64) Rectangle {
	return Rectangle{Width: r.Width * factor, Height: r.Height * factor}
}

// Non-mutator method - receives a COPY of the struct value, so does not modify the original instance.
// Area calculates and returns the area of the rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Create a new Rectangle instance with the specified width and height
// Consistency: If the type has mutator methods that need to modify the original instance,
// using a pointer consistently aligns with the mutability requirement.
func NewRectangle(w, h float64) *Rectangle {
	return &Rectangle{Width: w, Height: h}
}
