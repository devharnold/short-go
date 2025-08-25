package main

import (
	"fmt"
	"strings"
)

func main() {
	result := strings.HasPrefix("test", "te")
	fmt.Println("HasPrefix: ", result)
}