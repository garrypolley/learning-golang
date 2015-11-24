package main

// The use of time.Sleep is just for fun. It mostly shows
// how the queue is able to get values and take them off
// fairly quickly. It can also cause a deadlock though
// because there is a max size of 40 on the queue channel.
// Whenever a queue is at max size and a write is attempted it
// will dead lock.

import (
	"fmt"
	"time"
)

func main() {
	queue := make(chan string, 40)

	// Uncomment this section to see the dead lock error caused by
	// writing to an already filled up channel.
	// --> fatal error: all goroutines are asleep - deadlock!
	// for j := 0; j < 41; j++ {
	// 	queue <- "Value"
	// }

	go func() {
		for i := 0; i < 900000000000; i++ {
			time.Sleep(time.Second / (time.Second * 5))
			queue <- fmt.Sprintf("Made up number value %d", i)
		}
	}()

	go func() {
		queue <- "one"
		queue <- "two"
	}()

	// Generally it seems you close the queue if you know you will not
	// be putting anything else onto the queue later.
	// Also, it seems that if you do not close the queue or have the queue
	// operated on in another go routine you'll get a deadlock error.
	//
	// Better details found here: http://guzalexander.com/2013/12/06/golang-channels-tutorial.html
	// Basically, it's becuase the write will not happen on a channel until their is a listener.
	// There won't be a listener without a goroutine since the code execution is synchronous

	i := 0
	for elem := range queue {
		time.Sleep(time.Second / (time.Second * 6))
		fmt.Println(elem)
		fmt.Println("Queue length: ", len(queue))
		if elem == "two" {
			fmt.Println()
			fmt.Println("iteration: ", i)
			fmt.Println()
			i++
			queue <- "one"
			queue <- "two"
			fmt.Println("added more to the chain")
			if i > 40 {
				break
			}
		}
	}

	fmt.Println("items on queue for loop")
	for i := 0; i < len(queue); i++ {
		fmt.Println("Item from queue: ", <-queue)
	}

}
