// Each Go source file begins with a package declaration
package main

// import an office package into the file, we use the Import instruction in the following way:

// Alternatively, you can use parentheses for multiple imports
import (
	"fmt"
	"playground/basics"
)

func raining() bool {
	fmt.Println("Check if it is raining now...")
	return true
}

func snowing() bool {
	fmt.Println("Check if it is snowing now...")
	return true
}

func shortCircuitBasics() {

	if raining() || snowing() {
		fmt.Println("Stay indoors!")
	}
}

func main() {
	fmt.Println("Hello, there!")

	// shortCircuitBasics()

	// basics.VariableBasics()
	// basics.InputOutputBasics()

	// basics.BitShiftingBasics()
	// basics.SwitchBasics()
	// basics.DisplayTime()
	// basics.LoopBasics()
	basics.StructBasics()

	// basics.PointerBasics()
	// basics.FunctionBasics()
	// basics.ArrayBasics()
	// basics.SliceBasics()
	// basics.MapBasics()
	// basics.InterfaceBasics()
	// basics.MethodBasics()
	// basics.DeferBasics()

	// exercises.NumberGuessingGame()
	// exercises.ScoreManager()

}
