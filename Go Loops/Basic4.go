package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"strings"
	"sync"
	"time"
)

func lower_bound(array []int, target int) int {
	left, right := 0, len(array)

	for (right - left) > 0 {
		mid := left + (right-left)/2

		if array[mid] < target {
			left = mid + 1

		} else {
			right = mid
		}

	}

	return left
}

var hashArray []string

func generateHash(size int, casing string, mtx *sync.Mutex, wg *sync.WaitGroup) {
	charset := "abcdefghijklmnopqrstuvwzyz1234567890"
	var builder strings.Builder // for better string concatination

	for i := 0; i < size; i++ {
		index, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset)))) // using the inbult crypto.rand method for better seeding
		if err != nil {
			log.Fatal(err)
		}

		builder.WriteByte(charset[index.Int64()])
	}

	mtx.Lock()
	if casing == "upper" {
		hashArray = append(hashArray, strings.ToUpper(builder.String()))
	} else {
		hashArray = append(hashArray, builder.String())
	}
	mtx.Unlock()

	wg.Done()

}

func main() {
	array := []int{51, 55, 72, 80, 100, 29, 40}

	mtx := &sync.Mutex{}
	wg := &sync.WaitGroup{}

	// Sort the array in accending order.
	// sort.Ints(array) //

	// fmt.Println(array)
	// target := 5
	// fmt.Printf("lower bound of %v is %v", target, lower_bound(array, target))

	start := time.Now()

	for _, v := range array {
		wg.Add(1)
		go generateHash(v, "upper", mtx, wg)
		// fmt.Printf("Hash of size %v is generated: %v", v, hashcode)
	}

	wg.Wait()

	elapsed := time.Since(start)

	fmt.Printf("All go-routines are finished at %#v\n", elapsed.Milliseconds())
	for _, v := range hashArray {
		fmt.Printf("Hash of size %v is generated: %v\n", len(v), v)
	}

}
