package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker ", id, " procesing job ", j)
		time.Sleep(time.Second) // Sleep one second to similuate work being done
		results <- j * 2
	}
}

func main() {

	// Play wth these two numbers to see how fast something can get done.
	// The sleep similutes a long running job, so this right now takes about 5 seconds
	// even though there are 100 jobs to do.
	const NUM_JOBS = 1000
	const NUM_WORKERS = 200

	jobs := make(chan int, NUM_JOBS)
	results := make(chan int, NUM_JOBS)

	for w := 1; w <= NUM_WORKERS; w++ {
		go worker(w, jobs, results)
	}

	for w := 1; w <= NUM_JOBS; w++ {
		jobs <- w
	}
	close(jobs) // Says we are done doing jobs

	for w := 1; w <= NUM_JOBS; w++ {
		<-results // Get all the values off of results that were added
	}
}
