// Each Go source file begins with a package declaration
package main

// import an office package into the file, we use the Import instruction in the following way:

// Alternatively, you can use parentheses for multiple imports
import (
	"fmt"
	"playground/basics"
	"reflect"
)

func raining() bool {
    fmt.Println("Check if it is raining now...")
    return true
}

func snowing() bool {
    fmt.Println("Check if it is snowing now...")
    return true
}

func shortCircuitBasics(){

    if raining() || snowing() {
        fmt.Println("Stay indoors!")
    }    
}

func variables(){
    // declare and initialize variable using the var keyword followed by the variable name and the type
    // Specifying the data type: Explicitly typed variable
    var score int = 75

    // declared as float32 and then initialized
    var rate float32 = 4.5 
    
    // Go can infer the type if you provide an initial value:
    var isActive = true

    // You can declare multiple variables in a single statement
    // Constants are variables whose values cannot be changed once they are set
    const name, age = "MK", 42

    // Short Variable Decleration Operator :=
    // Declare firstName as a string variable by initializing it, all without needing to use the var prefix.
    firstName := "Wei-Meng"

    // reflect package allows you to find out the data type of a variable
    fmt.Println(reflect.TypeOf(firstName)) // string

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
    isEven := num % 2 == 0

    fmt.Println(num, isInRange, isEven) // 11 false false

    if isEven {
        fmt.Printf("%d is even.\n", num)
    } else {
        fmt.Printf("%d is odd.\n", num)
    }


}



func main() {
    fmt.Println("Hello, there!")
    // variables()
    // shortCircuitBasics()

    // basics.DisplayTime()
    // basics.LoopBasics()
    basics.StructBasics()


}


