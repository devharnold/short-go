package main

import "fmt"

func sumNumbers(numbers ...int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func main() {
	total := sumNumbers(1, 2, 3, 4)
	fmt.Println("Total: ", total)
}