package main

import (
	"fmt"
	"learning_go/try-imports/constants" // This must be on your go path to compile
)

func printLoop(options []string) {
	for _, elem := range options {
		fmt.Println("Constant value is: ", elem)
	}
}

func main() {
	fmt.Println("trying this package out")

	printLoop(constants.Options)

	// Copy here is what causes the slice to not get overridden.
	// Also the length is needed. Copy copies the elements based on
	// the smallest size of the slices.
	newOptions := make([]string, len(constants.Options))
	copy(newOptions, constants.Options)

	newOptions[0] = "No breakage made"

	printLoop(constants.Options)

	// Example of breaking the constants
	breakingOptions := constants.Options

	breakingOptions[0] = "busted that up"
	printLoop(constants.Options)
}
