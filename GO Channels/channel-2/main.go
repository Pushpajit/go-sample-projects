package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"sync"
	"time"
)

var (
	array  []int
	record time.Time
)

func recordStart() {
	record = time.Now()
}

func recordEnd() {
	endtime := time.Since(record)
	fmt.Printf("Time took: %#+v ms\n", endtime.Milliseconds())
}

func main() {
	wg := &sync.WaitGroup{}
	mtx := &sync.Mutex{}

	// predefining the variables
	totalElements := 40000000
	numGoroutines := 5
	array = make([]int, 0, totalElements)

	// starting the go-routine
	recordStart()
	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go generateRandomArray2(totalElements/numGoroutines, wg, mtx, fmt.Sprintf("thread-%d", i+1))
	}
	wg.Wait()

	fmt.Printf("Array of size %v is created\n", len(array))
	recordEnd()

	// searching into the array
	var target int
	fmt.Println("Enter a number you want to search: ")
	fmt.Scanf("%v", &target)

	recordStart()
	for _, v := range array {
		if v == target {
			fmt.Printf("%v is found in the array\n", target)
			break
		}
	}
	recordEnd()

}

// This is un-optimized version.
func generateRandomArray(n int, wg *sync.WaitGroup, mtx *sync.Mutex, thread string) {
	defer wg.Done()

	for i := 0; i < n; i++ {
		ranNum, err := rand.Int(rand.Reader, big.NewInt(int64(n)))
		if err != nil {
			log.Panic(err)
		}
		mtx.Lock()
		array = append(array, int(ranNum.Int64())) // This is costly method.
		mtx.Unlock()
	}

	fmt.Println(thread + " finished.")

}

// This is optimized version.
func generateRandomArray2(n int, wg *sync.WaitGroup, mtx *sync.Mutex, thread string) {
	defer wg.Done()

	// make a local array, insead of directly changing the global array.
	localArray := make([]int, n)

	for i := 0; i < n; i++ {
		ranNum, err := rand.Int(rand.Reader, big.NewInt(int64(n)))
		if err != nil {
			log.Panic(err)
		}
		localArray[i] = int(ranNum.Int64())
	}

	mtx.Lock()
	array = append(array, localArray...) // Append in one go
	mtx.Unlock()

	fmt.Println(thread + " finished.")
}
