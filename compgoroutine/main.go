// illustrating complex Goroutine where multiple tasks communicate

package main

import (
	"fmt"
	"time"
)

// jobs <-chan: This is a receive only channel of `int`, the workers read jobs from it
// results chan<-: This is a send-only channel of `int`, the workers write results onto it

func worker(id int, jobs <-chan int, results chan<-int) {
	for j := range jobs {
		fmt.Println("Worker", id, "processing job", j)
		time.Sleep(time.Second)
		results <- j * 2
	}
}

func main() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= numJobs; a++ {
		fmt.Println("Result", <-results)
	}
}