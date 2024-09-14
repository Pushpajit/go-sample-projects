package main

import (
	"fmt"
	"sync"
)

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 30, 23, 43, 14}
	target := 10

	pipe := make(chan bool, 1)
	wg := &sync.WaitGroup{}

	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < len(array)/2; i++ {
			if target == array[i] {
				select {
				case pipe <- true:
					fmt.Println("The target is found by 1st go-routine")
				default:
				}
				return
			}
		}
	}()

	go func() {
		defer wg.Done()
		for i := len(array) / 2; i < len(array); i++ {
			if target == array[i] {
				select {
				case pipe <- true:
					fmt.Println("The target is found by 2nd go-routine")
				default:
				}
				return
			}
		}
	}()

	go func() {
		wg.Wait()
		close(pipe)
	}()

	<-pipe
	fmt.Println("All go-routine finished")
}
