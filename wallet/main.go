package main

import "fmt"

func main() {
	appName := "Wallet"
	version := 1
	user := "Henry"
	balance := 100.50

	var errorCount int
	if balance < 0 {
		errorCount++
	}

	fmt.Println("App:", appName)
	fmt.Println("version:", version)
	fmt.Println("User:", user)
	fmt.Println("Balance:", balance)
	fmt.Println("Errors:", errorCount)
}