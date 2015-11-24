package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Millisecond * 500)
	go func() {
		for t := range ticker.C {
			fmt.Println("tick is at: ", t)
		}
	}()

	// These two lines are basically a sleep statement.
	// time.Sleep(time.Second * 4)
	timer := time.NewTimer(time.Second * 4)
	<-timer.C

	ticker.Stop()
	fmt.Println("Ticker Stopped")
}
