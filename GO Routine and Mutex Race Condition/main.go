package main

import (
	"fmt"
	"sync"
)

// declearing a global resouce.
var resource []int

func main() {
	// creating a waitgroup for thread sync.
	wg := &sync.WaitGroup{}
	// creating mutex for avoiding deadlock.
	mtx := &sync.Mutex{}

	for i := 1; i <= 10; i++ {
		// adding go-routines into the waitgroup one at a time.
		wg.Add(1)

		// forking the function into a seperate go-routine.
		go job(i, wg, mtx)
	}

	// waiting for the go-routines to finish their task.
	wg.Wait()

	// printing the final resource.
	fmt.Printf("%#v", resource) // remember go-routine never maintain the correct order of their execution.
}

func job(jobid int, wg *sync.WaitGroup, mtx *sync.Mutex) {
	// after finishing the task it will signal the main go-routines (main thread) that the job is done.
	defer wg.Done()

	// printing the currently running routine.
	fmt.Printf("job %v is running\n", jobid)

	// Avoding Race-Condition.
	// So no two threads can write in the same memory address at the same time.
	mtx.Lock()                         // locking resource.
	resource = append(resource, jobid) // writting into the resource.
	mtx.Unlock()                       // unlocking resource.
}
