package basics

import (
	"fmt"
	"log"
)

func LogBasics() {

	fmt.Println("Logging Basics")
	fmt.Println("--------------")

	// Logging messages
	log.Print("This is a log message.")                   // 2024/01/04 09:33:36 This is a log message.
	log.Printf("Formatted log message: %s", "Hello, Go!") // 2024/01/04 09:33:36 Formatted log message: Hello, Go!

	// Fatal logs a message and exits the program with a non-zero status code
	// Fatal is equivalent to l.Print() followed by a call to os.Exit(1)
	log.Fatal("Fatal error occurred.") // exit status 1

	fmt.Println("This line will NOT be printed")

}
