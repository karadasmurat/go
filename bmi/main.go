package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// Both bufio.Scanner and fmt.Scanf are valid options for getting a single float from user input in Go
// For simplicity and brevity when only needing a single float, fmt.Scanf might be more efficient.
var reader = bufio.NewReader(os.Stdin)

func main() {
	fmt.Println("Welcome to the BMI Calculator")
	fmt.Println("-----------------------------")

	weight, err := getFloatInput("Enter weight (kg): ")
	if err != nil {
		log.Fatal(err)
	}

	height, err := getFloatInput("Enter height (m): ")
	if err != nil {
		log.Fatal(err)
	}

	bmi_val := bmi(weight, height)

	fmt.Println("BMI: ", bmi_val)

	interpretBMI(bmi_val)

}

func getFloatInput(msg string) (float64, error) {
	var num float64
	fmt.Print(msg)

	_, err := fmt.Scanf("%f", &num)

	if err != nil {
		return 0, fmt.Errorf("invalid input: %w", err)
	}
	if num <= 0 {
		return 0, fmt.Errorf("Invalid Input: Positive numbers only.")
	}
	return num, nil
}

func bmi(w float64, h float64) float64 {
	return w / (h * h)
}

func interpretBMI(bmi float64) {
	var category string
	if bmi < 18.5 {
		category = "Underweight"
	} else if bmi < 25 {
		category = "Normal weight"
	} else if bmi < 30 {
		category = "Overweight"
	} else {
		category = "Obese"
	}
	fmt.Printf("BMI category: %s\n", category)
}
