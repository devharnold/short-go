package main

import (
	"fmt"
	"short-go/utils"
)

func main() {
	utils.PrintHeader("Maps Demo")

	// creating a map in Go
	userAges := map[string]int{
		"Henry": 25,
		"Alice": 30,
	}

	// adding a new key-value pair to the map
	userAges["Bob"] = 20

	// access a value
	fmt.Println("Henry's age:", userAges["Henry"])

	// checking if a key exists
	age, exists := userAges["Eve"]
	if exists {
		fmt.Println("Eve's age:", age)
	} else {
		fmt.Println("Eve not found in map")
	}

	// iterate over a map
	fmt.Println("All Users:")
	for name, age := range userAges {
		fmt.Println("%s -> %d\n", name, age)
	}
}