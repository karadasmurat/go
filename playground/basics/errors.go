package basics

import (
	"errors"
	"fmt"
	"log"
)

type MKError struct {
    Message string
}

func (e *MKError) Error() string {
    return e.Message
}

func returnCustomError() error {
    return &MKError{Message: "This is a custom error"}
}


func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func ErrorBasics() {
	fmt.Println("Error Basics")
	fmt.Println("------------")

	// When calling a function that returns an error, it's common to check for errors immediately.
	res, err := divide(10, 0)
	if err != nil {
		log.Fatal("Division by zero.") // exit status 1
	}
	fmt.Println(res)
}
