package model

// Author struct represents the author of a book
type Author struct {
	Name string
}

// Book struct represents a book with an author using a pointer
type Book struct {
	Title  string
	Author *Author // Pointer member to the Author
}
