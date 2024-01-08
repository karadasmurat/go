package basics

import "fmt"

type ImmutableStruct struct {
	Count  int
	Status string
}

func seemsToModify(is ImmutableStruct) ImmutableStruct {

	// Return a new struct with modifications
	return ImmutableStruct{
		Count:  is.Count + 1,
		Status: "Modified",
	}

}

// method on immutable type - avoiding operations that would modify receiver
// Instead, create and return new instances.
func (s ImmutableStruct) AddOne() ImmutableStruct {
	return ImmutableStruct{
		Count:  s.Count + 1,
		Status: "Incremented",
	}
}

func ImmutabilityBasics() {
	fmt.Println("Immutability Basics")
	fmt.Println("-------------------")

	is1 := ImmutableStruct{Count: 0, Status: "Init"}
	fmt.Printf("is1@ %p :%v\n", &is1, is1) // is1@ 0x140000b6000 :{0 Init}

	// seems to modify is1, but it does not!
	is1.AddOne()
	fmt.Printf("is1@ %p :%v\n", &is1, is1) // is1@ 0x140000b6000 :{0 Init}

	is2 := is1.AddOne()
	fmt.Printf("is1@ %p :%v\n", &is1, is1) // is1@ 0x140000b6000 :{0 Init}
	fmt.Printf("is2@ %p :%v\n", &is2, is2) // is2@ 0x140000b6048 :{1 Incremented}
}
