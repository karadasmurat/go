package basics

import "fmt"

func FunctionBasics() {

	fmt.Println("Function Basics")
	fmt.Println("---------------")

	// Variables of a function type
	// Declare i to be a function that returns an int value:
	var i func() int
	i = giveMeZero

	// calling a variable of a function type
	val := i()
	fmt.Println(val)

	// Anonymous function definition and assignment:
	// func() { ... } defines an anonymous function without a name.
	// f := assigns the function to the variable f, allowing you to refer to it later.
	f := func() {
		fmt.Println("Hello, this is an anonymous function.")
	}

	// Call the function stored in the variable f
	f()

	// Define an anonymous function and immediately invoke it
	func() {
		fmt.Println("Hello, immediately invoked anonymous!")
	}()

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
