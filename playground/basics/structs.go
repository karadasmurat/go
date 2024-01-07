package basics

import (
	"fmt"
	"playground/model"
	"reflect"
	"unsafe"
)

type Opts struct {
	configItem1 int
	configItem2 string
	configItem3 bool
	timeout     int
	tls         bool
}

func (o *Opts) withTLS() *Opts {
	o.tls = true // Pointer receiver - allowing the method to modify the original instance.
	return o
}

func NewOpts() *Opts {
	// Can we do this on C? Passing the address of a local variable!
	return &Opts{
		configItem1: 42,
		configItem2: "Default",
		configItem3: true,
	}
}

type Server struct {
	Opts
}

// Constructor function that returns a pointer to a new instance of Server
func NewServer(opts *Opts) *Server {
	// Can we do this on C? Passing the address of a local variable!
	return &Server{
		Opts: *opts,
	}
}

func pointerToStruct() {

	fmt.Println("Pointer to Struct")
	fmt.Println("-----------------")

	potter := model.Wizard{Name: "Potter", House: "Gryffindoor"}

	fmt.Printf("%v @ %p\n", potter, &potter)             // {Potter Gryffindoor} @ 0x140000b6060
	fmt.Printf("%v @ %p\n", potter.Name, &potter.Name)   // Potter @ 0x140000b6060
	fmt.Printf("%v @ %p\n", potter.House, &potter.House) // Gryffindoor @ 0x140000b6070
}

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

	pointerToStruct()

	embeddedStruct()

	author1 := model.Author{Name: "author1"}
	fmt.Printf("author1@ %p\n", &author1) // author1@ 0x1400008e090
	book1 := model.Book{Title: "title01", Author: &author1}
	fmt.Printf("book1@ %p book1.Author= %p\n", &book1, book1.Author) // book1@ 0x140000c0018 book1.Author= 0x1400008e090

	fmt.Println("Before function call: ", book1.Title, book1.Author.Name) // Before function call:  title01 author1
	acceptStructContainingPointer(book1)
	fmt.Println("After function call: ", book1.Title, book1.Author.Name) // After function call:  title01 Modified Name

	configBasics()

	constructorsAndMutatorMethods()
}

func embeddedStruct() {

	contact := model.NewContact("MK", "IST", "TR")
	fmt.Println(contact.String()) // Name: MK, Address: IST, TR
}

// note that function accepts a struct, which has a pointer field. (like slice)
func acceptStructContainingPointer(b model.Book) {
	// b is a copy of the struct, at a new address
	// however, the value of b.Author is the same as the argument (the same value, at a different address)
	fmt.Printf("b@ %p b.author= %p\n", &b, b.Author) // b@ 0x140000c0030 b.author= 0x1400008e090

	// modify title (value) will NOT be visible
	b.Title = "Modified Title"

	// modify author (through pointer) WILL BE VISIBLE
	// b.Author.Name = "Modified Name" // b.Author->Name

	// modify pointer itself - will NOT be visible
	// b.Author = &model.Author{Name: "author2"}

	// what if we change the original as well?
	tmp := b.Author // 0x1400008e090
	b.Author = &model.Author{Name: "author2"}
	tmp.Name = "author2"

	// return the modified struct to make modifications visible
}

func constructorsAndMutatorMethods() {
	var c model.Counter = model.NewCounter(0)
	var cptr *model.Counter = model.NewCounterPtr(0)

	c.Increment()
	cptr.Increment() // Go allows implicit dereferencing

	fmt.Println(c.Count)
	fmt.Println(cptr.Count)
}

func configBasics() {

	// v1. Initializing a struct with a composite literal
	// This creates a new instance of Struct and initializes its fields directly.
	// var s1 Server = Server{maxConn: 1, id: "Server1", tls: false}
	// fmt.Println(s1) // {1 Server1 false}

	// v2. Using the constructor to create a new instance of Struct.
	// var s2 *Server = NewServer(1, "Server2", false)

	// Accessing the fields of the created struct.
	// Implicit dereferencing: Go automatically dereferences the pointer for you: (*s1).tls
	// fmt.Println(s2.tls)
	// fmt.Println((*s2).tls)

	// v3.a Using the constructor to create a new instance of Struct.
	// Empty config
	var s3 *Server = NewServer(&Opts{})
	fmt.Println(s3.configItem3)

	// v3.b Using a default config
	s4 := NewServer(NewOpts())
	fmt.Println(s4.tls)

	// v3.b Using a default config
	s5 := NewServer(NewOpts().withTLS())
	fmt.Println(s5.tls)

}
