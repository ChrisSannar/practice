package main

import (
	"fmt"
	"sync"
)

/*
* NOTE:
* - use `go run -race main.go` to detect race conditions
* - Limit scope of locks to minimize contention
* - Always unlock (use defer)
* - Consider using `RWMutex` when appropriate
* - Use buffered channels to reduce overhead
 */
func main() {
	fmt.Println("Race")
	basics()
	useMutex()
	useChannel()
}

func basics() {
	var value int
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			value++
		}()
	}

	wg.Wait()
	fmt.Println("Value:", value)
}

func useMutex() {
	var value int
	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			value++
			mu.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println("Value:", value)
}

func useChannel() {
	const numGoRoutines = 10

	dataCh := make(chan int, numGoRoutines)
	var wg sync.WaitGroup

	for i := 0; i < numGoRoutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			dataCh <- 1
		}()
	}

	go func() {
		wg.Wait()
		close(dataCh)
	}()

	total := 0
	for v := range dataCh {
		total += v
	}
	fmt.Println("Data value:", total)
}
