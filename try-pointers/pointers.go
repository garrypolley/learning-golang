package main

import (
	"fmt"
)

type Object struct {
	a int
	b int
	c int
}

// func (object Object) String() string {
// 	return fmt.Sprintf("\n{\n  a: %d,\n  b: %d,\n  c: %d\n}", object.a, object.b, object.c)
// }

func printValues(objects []*Object) {
	fmt.Println("Slice References are:")
	for _, object := range objects {
		fmt.Println(*object)
	}
	fmt.Println()
}

func main() {
	slice := []*Object{
		&Object{
			a: 1,
			b: 2,
			c: 3,
		},
		&Object{
			a: 4,
			b: 5,
			c: 6,
		},
		&Object{
			a: 7,
			b: 8,
			c: 9,
		},
	}

	fmt.Println("Starting Values")
	fmt.Println("Slice is: ", slice)
	printValues(slice)
}
