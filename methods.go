package main

import "fmt"

type rect struct {
	width  int
	height int
}

// Note to self, go will automatically do pointer reference
// for you on a method call. Which is pretty cool, though this
// can allow mutation, which isn't cool.
func (r *rect) area() int {
	return r.width * r.height
}

// This is not memory efficient because the rect is copied and
// passed by value instaed of reference. Go defaults to passing
// by value unless you tell it to use the reference.
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{
		width:  10,
		height: 5,
	}

	fmt.Println("area: ", r.area())
	fmt.Println("perim: ", r.perim())

	// Not sure why I'd bother using a pointer here, but it
	// was in the example so I gave it a whirl.
	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim: ", rp.perim())

}
