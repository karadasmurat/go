package basics

import (
	"fmt"
	"playground/collections"
	"reflect"
)

type Rectangle struct {
	Width  float64
	Height float64
}

// Method with an implicit receiver parameter (r Rectangle)
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func VariableBasics() {

	// declare and initialize variable using the var keyword followed by the variable name and the type
	// Specifying the data type: Explicitly typed variable
	var score int = 75

	// declare as float64 and then initialize
	var rate float64 = 4.5

	// Type Inference
	// Go can infer the type if you provide an initial value:
	var isActive = true

	fmt.Printf("isActive: %t\n", isActive) // isActive: true

	// You can declare multiple variables in a single statement
	// Constants are variables whose values cannot be changed once they are set
	const name, age = "MK", 42

	// Short Variable Decleration Operator :=
	// Declare firstName as a string variable by initializing it, all without needing to use the var prefix.
	firstName := "Wei-Meng"

	// reflect package allows you to find out the data type of a variable
	fmt.Println(reflect.TypeOf(firstName))       // string
	fmt.Printf("%T, %v\n", firstName, firstName) // string, Wei-Meng

	fmt.Printf("%s %d %f\n", firstName, score, rate)

	// %t to format a boolean as true or false.
	fmt.Printf("Active: %t\n", isActive)
	fmt.Printf("%s is %d years old.\n", name, age)

	// declares a variable named queue and assigns it the value 5.
	// The type of queue is inferred to be int because of the provided value.
	queue := 5
	lastname := "John"

	// Use fmt.Sprintf function to format a string.
	// The result is a formatted string stored in the variable s.
	s := fmt.Sprintf("%s, your queue number is %d", lastname, queue)

	fmt.Println(s)

	num := 12

	// Declare a variable isBigEnough and assigns it the result of the boolean expression
	isInRange := num > 1 && num < 10

	// Declare a variable isEven and assigns it the result of the boolean expression num % 2 == 0
	isEven := num%2 == 0

	fmt.Println(num, isInRange, isEven) // 11 false false

	if isEven {
		fmt.Printf("%d is even.\n", num)
	} else {
		fmt.Printf("%d is odd.\n", num)
	}

	// Raw Strings
	quotation := `"Anyone who has never made 
a mistake has never tried 
anything new."
		--Albert \nEinstein`

	fmt.Println(quotation)

}

func InputOutputBasics() {
	var name string
	var age int
	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)
	fmt.Print("Enter your age: ")
	fmt.Scanln(&age)
	fmt.Printf("Name: %s, Age: %d\n", name, age)
}

func BitShiftingBasics() {
	fmt.Println("Bit Shifting Basics")
	fmt.Println("-------------------")

	n := 32 // 2^5

	n1 := n >> 3 // 2^5 / 2^3 = 2^2 = 4
	n2 := n << 3 // 2^5 * 2^3 = 2^8 = 256

	fmt.Println(n, n1, n2) // 32 4 256
}

func getDayString(n uint8) string {

	dayOfWeek := ""

	switch n {
	case 1:
		dayOfWeek = "Monday"
	case 2:
		dayOfWeek = "Tuesday"
	case 3:
		dayOfWeek = "Wednesday"
	case 4:
		dayOfWeek = "Thursday"
	case 5:
		dayOfWeek = "Friday"
	case 6:
		dayOfWeek = "Saturday"
	case 7:
		dayOfWeek = "Sunday"
	default:
		dayOfWeek = "--error--"
	}

	return dayOfWeek

}

func PassOrFail(grade string) string {
	result := "Undefined"
	switch grade {
	case "A", "B", "C", "D": // case statement to match multiple values
		result = "Passed"
	case "F":
		result = "Failed"
	default:
		result = "Undefined"
	}

	return result
}

func SwitchBasics() {
	daystr := getDayString(3)
	fmt.Print(daystr)

	result := PassOrFail("F")
	fmt.Print(result)
}

func LoopBasics() {

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	scores := [3]int{75, 90, 66}
	for index, value := range scores {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// If you don’t care about the index, you can use a blank identifier _
	for _, v := range scores {
		fmt.Println(v)
	}

	// Iterating through a string
	fmt.Println("Iterating through a string")
	for pos, char := range "Hello, world!" {
		fmt.Println(pos, char)           // 0 72
		fmt.Printf("%d %c\n", pos, char) // 0 H
	}

}

type Vertex struct {
	X int
	Y int
}

func PointerBasics() {

	fmt.Println("Pointer Basics")
	fmt.Println("--------------")

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

// When you pass an array to a function, a copy of the entire array is passed.
func tryToModifyInputArray(arr [3]int) {

	arr[0] = 0
	fmt.Println("Inside function:", arr) // [0 ... ]
}

func modifyArray(arr *[3]int) {
	arr[0] = 0
}

func ArrayBasics() {
	fmt.Println("Array Basics")
	fmt.Println("------------")

	// Declaring an array called scores with a fixed size of five elements
	var scores [5]float32

	// Note that an array's length is part of its type
	// reflect.TypeOf(scores) vs %T
	fmt.Printf("%T %v\n", scores, scores) // [5]float32 [0 0 0 0 0]
	// Get the length of an array
	fmt.Println(len(scores)) // 5

	// Declare and initialize using an Array Literal
	grades := [5]string{"A", "B", "C", "D", "F"}
	fmt.Println(reflect.TypeOf(grades)) // [5]string
	fmt.Println(grades)                 // [A B C D F]

	// The index (the position of the elements in the array) of elements are zero-based
	// Access the first element in an array
	scores[0] = 90
	fmt.Println(scores) // [90 0 0 0 0]

	// copy arrays (not references)
	nums1 := [5]int{1, 2, 3, 4, 5}

	// Create a shallow copy of the array.
	numsCopy := nums1

	// modifications on the copy
	numsCopy[0] = 0

	fmt.Println(nums1)    // [1 2 3 4 5]
	fmt.Println(numsCopy) // [0 2 3 4 5]

	// When you pass an array to a function, a copy of the entire array is passed.
	nums2 := [3]int{11, 22, 33}
	tryToModifyInputArray(nums2)

	// Any modifications made to the array inside the function are only local to that function and do not affect the original array.
	fmt.Println(nums2) // Not modified! [11 22 33]

	modifyArray(&nums2)
	fmt.Println(nums2) // Modified! [0 22 33]
}

func SliceBasics() {
	fmt.Println("Slice Basics")
	fmt.Println("------------")

	// declared without an explicit initial value.
	// The zero value for a slice is nil, and when printed, it will be an empty slice []
	var s1 []int // nil

	// Empty slice of integers, using Slice Literal
	s2 := []int{}

	s3 := []string{"Apple", "Orange", "Banana", "Grape"}

	// Creating a slice with a length of 0  (no specified initial capacity)
	s4 := make([]int, 0)

	// Creating a slice with a length of 0  (no specified initial capacity)
	s5 := make([]int, 0, 10)

	fmt.Println(s1, len(s1), cap(s1), s1 == nil) // [] 0 0 true
	fmt.Println(s2, len(s2), cap(s2), s2 == nil) // [] 0 0 false
	fmt.Println(s3, len(s3), cap(s3), s3 == nil) // [Apple Orange Banana Grape] 4 4 false
	fmt.Println(s4, len(s4), cap(s4), s4 == nil) // [] 0 0 false
	fmt.Println(s5, len(s5), cap(s5), s5 == nil) // [] 0 10 false

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

	// Searching
	numbers := []int{-1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	q := 1
	if i, ok := collections.Contains(numbers, q); ok {
		fmt.Println(q, "found at index:", i)
	} else {
		fmt.Println("Not Found", q)
	}
}

func MapBasics() {
	fmt.Println("Map Basics")
	fmt.Println("------------")

	// By default, the map variable (heights) is pointing to nil before initialization using the make()
	var ages map[string]int

	fmt.Println(reflect.TypeOf(ages)) // map[string]int

	//  initialize it using the make() function before using
	grades := make(map[string]int)
	grades["Potter"] = 80
	grades["Weasley"] = 70
	grades["Hermione"] = 100

	// The len function is used to get the number of key-value pairs in a map:
	fmt.Printf("Number of key-value pairs: %v\n", len(grades)) // 3

	// Deleting Values
	delete(grades, "Weasley")

	// Accessing Values:
	gradeP := grades["Potter"]

	// Checking if a Key Exists:
	q := "Weasley"
	value, exists := grades[q]
	if exists {
		fmt.Printf("Key Found! grades[%v]: %v\n", q, value)
	} else {
		fmt.Println("Key not found.")
	}

	fmt.Println(grades, gradeP) // map[Hermione:100 Potter:80] 80

	// Initializing a map with a map literal
	heights := map[string]int{"Potter": 170, "Weasley": 168}

	fmt.Println(heights) // map[Potter:170 Weasley:168]

	// Iterating over a map
	for key, value := range heights {
		fmt.Printf("heights[%v]: %v\n", key, value)
	}

}

func DeferBasics() {
	fmt.Println("counting")

	for i := 0; i < 5; i++ {
		// A deferred function’s arguments are evaluated when the defer statement is evaluated.
		// Deferred function calls are pushed onto a stack. (executed in LIFO order after the surrounding function returns)
		defer fmt.Print(i)
	}

	fmt.Println("done")
}

type MyInterface interface {
	Method1() int
	Method2(int) (int, bool)
	// ... additional methods
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
