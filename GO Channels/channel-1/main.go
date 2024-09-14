package main

import (
	"fmt"
	"sync"
)

func main() {
	pipe := make(chan int)  // creating channel
	wg := &sync.WaitGroup{} // creating waitgroup for sync

	wg.Add(2) // initializing the thread-pool

	// channels need to send and consume data concurrently.
	go func(wg *sync.WaitGroup) { // fork-1 (sending data into the channel)
		pipe <- 10 // sending data to the channel
		wg.Done()  // this go-routine send a finished signal to main thread.
	}(wg)
	go func(wg *sync.WaitGroup) { // fork-2 (consuming data from the channel)
		fmt.Println(<-pipe) // consuming channel data
		wg.Done()           // this go-routine send a finished signal to main thread.
	}(wg)

	wg.Wait() // waiting for all go-routines to end their task and report back to main thread.

}
