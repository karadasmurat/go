package basics

import (
	"fmt"
	"playground/model"
)

func VariableBasics(){

	const name, age = "Kim", 22

	fmt.Printf("%s is %d years old.\n", name, age)
}

func LoopBasics(){
	
	for i:=0; i<5; i++ {
		fmt.Println(i)
	}

}

func StructBasics(){
	// Creating an instance of the Wizard struct
	var wiz1 model.Wizard = model.Wizard{Name: "Potter", House: "Gryffindoor"}

	// Creating an instance of the Wizard struct
    potter := model.Wizard{
        Name: "Harry Potter",
        House: "Gryffindor",
    }

	fmt.Println(wiz1)
	fmt.Println(potter.House)
}

