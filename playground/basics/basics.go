package basics

import (
	"fmt"
	"playground/model"
	"reflect"
)

func VariableBasics() {

	const name, age = "Kim", 22

	fmt.Printf("%s is %d years old.\n", name, age)
}

func LoopBasics() {

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

}

func StructBasics() {
	// v1 Creating an instance of the Wizard struct
	var potter model.Wizard = model.Wizard{Name: "Potter", House: "Gryffindoor"}

	// v2 Creating an instance of the Wizard struct
	weasley := model.Wizard{
		Name:  "Harry Potter",
		House: "Gryffindor",
	}

	// v3 Creating an instance of the Wizard struct
	hermione := model.Wizard{"Hermione Granger", "Gryffindor"}

	fmt.Println(potter)

	// Struct fields are accessed using a dot.
	fmt.Println(weasley.House, hermione.Name)

	v := Vertex{1, 2} // var v Vertex
	p := &v           // struct pointer
	p.X = 1e9
	fmt.Println(v)

	// 1. create a new instance of the Vertex struct
	// 2. The & operator is used to take the address of the newly created Vertex instance, resulting in a pointer to that instance.
	// 3. The resulting pointer is then assigned to the variable p2:
	p2 := &Vertex{1, 2} // has type *Vertex: var p2 *Vertex

	fmt.Printf("%p, %v", p2, *p2) // 0x14000110040, {1 2}%
}

type Vertex struct {
	X int
	Y int
}

func PointerBasics() {

	// Pointer decleration using the * symbol followed by the type of the variable it will point to.
	// The type *T is a pointer to a T value:
	var p1 *int
	fmt.Printf("%p\n", p1) // 0x0

	i := 42

	// The & operator generates a pointer to its operand.
	p2 := &i

	fmt.Printf("%v %p\n", i, &i) // 42 0x14000110018

	// "dereferencing" or "indirecting" through * operator:
	// read i through the pointer p2: *p2
	fmt.Printf("%p, %v\n", p2, *p2) // 0x14000110018, 42

	cnt := 1
	fmt.Println(cnt) // 1

	// the value of the argument (cnt) is passed to the function, therefore parameter x will have a copy of it
	increment_byval(cnt)
	fmt.Println(cnt) // 1

	cnt = increment_byval(cnt)
	fmt.Println(cnt) // 2

	increment_byptr(&cnt)
	fmt.Println(cnt) // 3

}

func ArrayBasics() {
	fmt.Println("Array Basics")
	fmt.Println("------------")

	// Declaring an array called scores with a fixed size of five elements
	var scores [5]float32

	// Note that an array's length is part of its type
	fmt.Println(reflect.TypeOf(scores)) // [5]float32
	// Get the length of an array
	fmt.Println(len(scores)) // 5
	fmt.Println(scores)      // [0 0 0 0 0]

	// Declare and initialize using an Array Literal
	grades := [5]string{"A", "B", "C", "D", "F"}
	fmt.Println(reflect.TypeOf(grades)) // [5]string
	fmt.Println(grades)                 // [A B C D F]

	// The index (the position of the elements in the array) of elements are zero-based
	// Access the first element in an array
	scores[0] = 90
	fmt.Println(scores) // [90 0 0 0 0]
}

func SliceBasics() {
	fmt.Println("Slice Basics")
	fmt.Println("------------")

	// Declare and initialize using a Slice Literal
	primes := []int{2, 3}

	fmt.Println(reflect.TypeOf(primes)) // []int
	// Find out the length and capacity of the slice
	fmt.Println(len(primes)) // 2
	fmt.Println(cap(primes)) // 2

	// Appending to a slice
	// when a slice has reached its capacity, appending more items cause all the existing items to be copied onto the new array,
	// together with the newly added items. The newly created underlying array will now have a bigger capacity.
	primes = append(primes, 5)

	fmt.Println(primes)      // [2 3 5]
	fmt.Println(len(primes)) // 3
	fmt.Println(cap(primes)) // 4

	names := [4]string{"Zero", "One", "Two", "Three"}
	fmt.Println(names) // [Zero One Two Three]

	// Select a half-open range which includes the first element, but excludes the last one:
	// names[low:high]
	a := names[0:2]                // [Zero One]
	fmt.Println(reflect.TypeOf(a)) // []string

	b := names[1:3] // [One Two]

	// If the starting index is 0, you can simply leave it out
	c := names[:2] // [Zero One]

	// to extract from a particular index through the end, you can leave out
	d := names[1:] // [One Two Three]

	fmt.Println(a, b, c, d)

	// Iterating through a slice
	//   Expected output:
	//   0 Zero
	//   1 One
	for index, value := range a {
		fmt.Println(index, value)
	}
}

func FunctionBasics() {
	// Variables of a function type
	// Declare i to be a function that returns an int value:
	var i func() int
	i = giveMeZero

	// calling a variable of a function type
	val := i()
	fmt.Println(val)

	// variable number of arguments
	fmt.Println(sum(1))
	fmt.Println(sum(1, 2))
	fmt.Println(sum(1, 2, 3))

	// closures
	fmt.Println(multiply(5, 10))

	// Create a closure where a is 5
	// call the outer function, which returns the inner function, thus call it as well
	// times(5)(10) or
	// multiplyBy5(10)
	multiplyBy5 := times(5)

	times100 := times(100)

	fmt.Println(times(5)(10))
	fmt.Println(multiplyBy5(10))
	fmt.Println(times100(2))
}

// Everything in Go is passed by value.
// In pass by value, the actual value of the argument is passed to the function.
// For example, in the below function, the parameter x is a copy of the value passed to it.
// The function works with a copy of the original value, and any modifications made to the parameter inside the function do not affect the original variable outside the function (argument).
func increment_byval(x int) int {
	return x + 1
}

// We want the address of the original argument to be able to modify it!
func increment_byptr(x *int) {
	*x++
}

// A variadic function takes in a variable number of arguments
// Essentially, nums is now a slice of int values:
func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

func giveMeZero() int {
	return 0
}

func multiply(a, b int) int {
	return a * b
}

// Closures
// times is a function that that takes an integer parameter, and
// returns an anonymous  function of type: func(int) int
// The Inner Function (Closure) takes an integer parameter b
// This returned function is a closure because it "closes over" the variable a, meaning it captures and remembers the value of a
func times(a int) func(int) int {
	return func(b int) int {
		return a * b // a is defined outside the function: closure
	}
}
