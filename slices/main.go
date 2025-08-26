package main

import (
	"fmt"
	"short-go/utils"
)

func main() {
	utils.PrintHeader("Slices Demo")

	// creating a slice
	numbers := []int{1, 2, 3, 4}
	fmt.Println("Initial slice:", numbers)

	// apend new elements
	numbers = append(numbers, 5, 6)
	fmt.Println("After append:", numbers)

	// slice a slice
	sub := numbers[1:4]
	fmt.Println("Sub-slice (1:4):", sub)

	// iterate over a slice
	fmt.Println("Iterating...:")
	for i, v := range numbers {
		fmt.Println("Index %d  -> value %d\n", i, v)
	}
}