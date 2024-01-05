package basics

import (
	"fmt"
	"playground/model"
	"reflect"
	"unsafe"
)

// Note that struct is a value type.
func tryToModifyStruct(rect Rectangle) {
	rect.Width = 0
}

func modifyStruct(rect *Rectangle) {
	rect.Width = 0
}

func StructBasics() {
	fmt.Println("Struct Basics")
	fmt.Println("-------------")

	// v0 Creating an instance of the Wizard struct
	var malfoy model.Wizard
	malfoy.Name = "Malfoy"
	malfoy.House = "Slytherin"

	// v1 declare and initialize using Struct Literal
	var potter model.Wizard = model.Wizard{Name: "Potter", House: "Gryffindoor"}

	// v2 short declare and initialize using Struct Literal
	weasley := model.Wizard{
		Name:  "Harry Potter",
		House: "Gryffindor",
	}

	// v3 Creating an instance of the Wizard struct
	// provide values for the fields in the order in which they are declared in the struct definition, or use field names as keys.
	hermione := model.Wizard{"Hermione Granger", "Gryffindor"}

	fmt.Println(potter)

	// Struct fields are accessed using a dot.
	fmt.Println(weasley.House, hermione.Name)

	// structs are value types
	weasleyTwin := weasley
	fmt.Printf("weasley@ %p\ncopy@ %p", &weasley, &weasleyTwin) // different addresses!

	// Declare an empty struct type
	// It takes up zero bytes of memory, and creating instances of it is extremely lightweight.
	var myEmptyStruct struct{}
	fmt.Printf("Size of empty struct: %d bytes\n", unsafe.Sizeof(myEmptyStruct)) // 0 bytes

	// Create an instance of the empty struct
	instance := struct{}{}
	fmt.Printf("Instance of empty struct: %v size: %v bytes\n", instance, unsafe.Sizeof(instance)) // {} 0 bytes

	// Anonymous struct definition and initialization
	wiz03 := struct {
		Age  int
		Name string
	}{
		Age:  10,
		Name: "Potter",
	}

	// Print the initialized struct
	fmt.Println("Anonymous Struct: ", wiz03) // {10 Potter}

	v := Vertex{1, 2} // var v Vertex
	p := &v           // struct pointer
	p.X = 1e9
	fmt.Println(v)

	// 1. create a new instance of the Vertex struct
	// 2. The & operator is used to take the address of the newly created Vertex instance, resulting in a pointer to that instance.
	// 3. The resulting pointer is then assigned to the variable p2:
	p2 := &Vertex{1, 2} // has type *Vertex: var p2 *Vertex

	fmt.Printf("%p, %v\n", p2, *p2) // 0x14000110040, {1 2}%

	// declare and initialize using Struct Literal
	car1 := model.Car{Make: "Kia", Model: "Sorento", Year: 2007}
	fmt.Println(reflect.TypeOf(car1)) // model.Car
	fmt.Println(car1)                 // {Kia Sorento 2007}

	myRect := Rectangle{Width: 100, Height: 50}

	//Pass by value
	// When you pass a struct to a function, a copy of the entire struct is passed to the function.
	tryToModifyStruct(myRect)
	fmt.Println(myRect) // Not Modified! {100 50}

	modifyStruct(&myRect)
	fmt.Println(myRect) // Modified! {0 50}

	embeddedStruct()
}

func embeddedStruct() {

	contact := model.NewContact("MK", "IST", "TR")
	fmt.Println(contact.String()) // Name: MK, Address: IST, TR
}
