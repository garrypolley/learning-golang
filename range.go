package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum: ", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index: ", i)
		}
	}

	kvs := map[string]string{"a": "army", "b": "bear"}
	for key, value := range kvs {
		fmt.Printf("%s -> %s \n", key, value)
	}

	for i, c := range "go😢ëł😩" {
		fmt.Println(i, c)
	}
}
