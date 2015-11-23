package main

import (
	"fmt"
)

// ...int is kind of like *args in python, but for a specific argument
func sum(nums ...int) {
	fmt.Println(nums)
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println("total: ", total)
}

func main() {

	sum(1, 2, 3, 4, 10)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4, 5}
	sum(nums...) // This is like a splat operator in Python
}
