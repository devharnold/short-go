package main

import "fmt"

// simulate a large struct
type BigData struct {
	Numbers[1000000] int
}

// function takes struct by value (makes a huge copy)
func processByValue(data BigData) {
	fmt.Println("Length (By value): ", len(data.Numbers))
}

// function takes struct by pointer (just address no big copy)
func processByPointer(data *BigData) {
	fmt.Println("Length by pointer: ", len(data.Numbers))
}

func main() {
	big := BigData{}
	fmt.Println("Calling with value...")
	processByValue(big)

	fmt.Println("Calling with pointer...")
	processByPointer(&big)
}