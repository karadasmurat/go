package exercises

import "testing"

func TestGetInitials(t *testing.T) {
	// call the function to be tested
	result := getInitials("Harry Potter")

	expected := "HP"

	if result != expected {
		t.Errorf("Returned %v, expected %v", result, expected)
	}
}
