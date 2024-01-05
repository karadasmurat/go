package model

import "fmt"

// Address struct represents a mailing address
type Address struct {
	City, Country string
}

// Contact struct represents a person with a mailing address
type Contact struct {
	Name    string
	Address // Embedded Address struct
}

// String returns a string representation of the Contact
func (c Contact) String() string {

	// We can access the embedded struct's fields directly
	// c.Name vs c.Address.Name
	return fmt.Sprintf("Name: %s, Address: %s, %s", c.Name, c.City, c.Country)
}

// constructor function for creating a new Contact instance.
func NewContact(name string, city string, country string) Contact {

	// When creating structs with literals, we have to initialize the embedding explicitly
	return Contact{Name: name, Address: Address{City: city, Country: country}}
}
