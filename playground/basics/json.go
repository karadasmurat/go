package basics

import (
	"encoding/json"
	"fmt"
	"playground/model"
)

func JsonBasics() {
	fmt.Println("JSON Basics")
	fmt.Println("-----------")

	potter := model.Wizard{Name: "Potter", House: "Gryffindor"}

	// Marshaling: Converting a Go struct to JSON
	// The resulting jsonData is a byte slice containing the JSON-formatted data
	jsonData, err := json.Marshal(potter)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(jsonData))

	// JSON data as a string
	jsonString := `{"name":"Malfoy", "house":"Slytherin"}`

	// variable to hold the unmarshaled data
	var wiz model.Wizard

	err = json.Unmarshal([]byte(jsonString), &wiz)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(wiz) // {Malfoy Slytherin}

}
