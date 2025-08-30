// synchronizing Goroutines
package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup // this counter will keep track of the number of Goroutines working

func worker(id int) {
	defer wg.Done()
	fmt.Println("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i)
	}
	wg.Wait()
}