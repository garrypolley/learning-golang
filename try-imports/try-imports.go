package main

import (
	"fmt"
	// This must be on your go path to compile
	// for me that meant having a go path of
	// GOPATH=/Users/garrypolley/gohome:/Users/garrypolley/
	// This allows me to play around with my local stuff
	// since I have a /Users/garrypolley/src folder
	// but this does not break any of my existing golang
	// work or projects since it looks in my
	// /Users/garrypolley/src last
	"learning_go/try-imports/constants"
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
