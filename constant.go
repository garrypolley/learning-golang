package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	fmt.Println(s)

	const bigInt = 439439439
	const n = 500000000

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(n)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
	fmt.Println(math.Sin(bigInt))
}
