package main

import (
	"fmt"
	"sync"
)

/*
* NOTE:
* - always call `wg.Done()` inside the routine
* - use `wg.Add()` to keep account
* - always pass pointers (`&wg`) avoid copying
* - make sure to use `wg.Wait()`
* - keep track of the count (else `wg.Wait()` will wait irractically)
* - minimize global variables
* - keep critical sections as small as possible
* - use channel based sync to handle complex communicaiton
* - use buffered channels on errors to prevent blocking
 */
func main() {
	fmt.Println("Wait Groups")
	// basics()
	withErrors()
}

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Processing task %d\n", id)

	fmt.Printf("Task %d completed\n", id)
}

func basics() {
	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait()
	fmt.Println("All tasks completed")
}

const numWorkers = 5

func worker2(id int, wg *sync.WaitGroup, errChan chan<- error) {
	defer wg.Done()

	if id%2 == 0 {
		errChan <- fmt.Errorf("task %d encountered an error", id)
		return
	}

	fmt.Printf("Task %d completed successfully\n", id)
}

func withErrors() {
	var wg sync.WaitGroup
	errChan := make(chan error, numWorkers)

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker2(i, &wg, errChan)
	}

	wg.Wait()
	close(errChan)

	// fmt.Println("Chan", errChan)

	for err := range errChan {
		if err != nil {
			fmt.Println("Error:", err)
		}
	}
	fmt.Println("All tasks completed")
}
