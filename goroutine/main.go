package main

import (
	"fmt"
	"time"
)

func printMessage() {
	fmt.Println("Hello from Goroutine")
}

func main() {
	go printMessage()
	fmt.Println("Hello from main function")
	// wait for the goroutine to finish
	time.Sleep(time.Second)
}