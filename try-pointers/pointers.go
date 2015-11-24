package main

import (
	"fmt"
)

type StructStuff struct {
	a int
}

// You can also implement the method for GetA
// on the embedded struct
// func (object StructStuff) GetA() int {
// 	return object.a
// }

type Object struct {
	StructStuff
	b int
	c int
}

// Since there is an embedded struct use the
// New{Struct} pattern so it is hidden from a
// consumer.
func NewObject(a int, b int, c int) *Object {
	return &Object{
		StructStuff: StructStuff{a: a},
		b:           b,
		c:           c,
	}
}

type Stuff interface {
	GetA() int
	GetB() int
}

func (object Object) GetA() int {
	return object.a
}

func (object Object) GetB() int {
	return object.b
}

// This is invalid and does not work.
// You cannot have a pointer to an interface
// func AllAEqualB(objects []*Stuff) bool {
// 	value := true

// 	for _, object := range objects {
// 		value = (*object).GetA() == (*object).GetB()
// 	}

// 	return value
// }

func AllAEqualB(objects []Stuff) bool {
	value := true

	for _, object := range objects {
		value = (object).GetA() == (object).GetB()
	}

	return value
}

func ConvertObjectSliceToStuff(objects []*Object) []Stuff {
	newStuff := make([]Stuff, len(objects))

	for index, data := range objects {
		newStuff[index] = data
	}

	return newStuff
}

func printValues(objects []*Object) {
	fmt.Println("Slice References are:")
	for _, object := range objects {
		fmt.Println(*object)
	}
	fmt.Println()
}

func main() {
	slice := []*Object{
		NewObject(1, 2, 3),
		NewObject(4, 5, 6),
		NewObject(7, 8, 9),
	}

	fmt.Println("Starting Values")
	fmt.Println("Slice is: ", slice)
	printValues(slice)

	commonSlice := []*Object{
		NewObject(1, 1, 3),
		NewObject(1, 1, 5),
		NewObject(1, 1, 7),
	}

	fmt.Println("Try out using some slice and references")

	fmt.Println("commonSlice is: ", commonSlice)

	stuffSlice := ConvertObjectSliceToStuff(commonSlice)

	fmt.Println("stuff slice is: ", stuffSlice)

	fmt.Println("All a == b in the commonSlice: ", AllAEqualB(stuffSlice))
}
