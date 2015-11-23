package main

// This deviates form the example on gobyexample.com
// because a fast machine will make it seem like the
// jobs channel is working synchronously when it is
// not acting that way. Adding the sleep helps illustrate
// the point of how the channels communicate and how they
// work when closing.

import (
	"fmt"
	"time"
)

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool, 1)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	go func() {
		for j := 1; j <= 3; j++ {
			jobs <- j
			fmt.Println("sent job")
			if j%2 == 0 {
				time.Sleep(time.Second * 1)
			}
		}
		close(jobs)
		// It's like the financial crisis
		fmt.Println("sent all jobs")
	}()

	<-done
}
