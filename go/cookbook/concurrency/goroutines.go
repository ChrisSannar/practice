package main

import (
	"fmt"
	"time"
)

/**
* NOTE:
* - Use Channels, wait groups, and mutexts for goroutines
* - For accessive internal comunication, used shared memory
* - Make sure they terminate
* - Make sure to use sync else you'll get race conditions
* - Watch channels for deadlocks
* - Handle uncaught panics with `recover`
* - Consider batching tasks to reduce synchronization
* - Load balance the routines
* - Use lazy initialization till routines are needed
* - Reduce context switching to reduce overhead
 */
func main() {
	fmt.Println("Goroutines")
	// basics()
	// anonFunc()
	parellel()
}

func processData() {
	for i := 0; i < 5; i++ {
		fmt.Printf("Processing data: %d\n", i)
		time.Sleep(1 * time.Second)
	}
}

func basics() {
	go processData()
	fmt.Println("Post goroutine")

	time.Sleep(6 * time.Second)
	fmt.Println("Processing complete")
}

func anonFunc() {
	go func() {
		for i := 'a'; i < 'e'; i++ {
			fmt.Printf("%c", i)
			time.Sleep(500 * time.Millisecond)
		}
	}()

	time.Sleep(3 * time.Second)
	fmt.Println("\nProcessing complete")
}

func worker(id int, work chan int, done chan bool) {
	for n := range work {
		fmt.Printf("Worker %d processing task %d\n", id, n)
		time.Sleep(time.Second)
		done <- true
	}
}

func parellel() {
	work := make(chan int, 5)
	done := make(chan bool, 5)

	for i := 1; i < 5; i++ {
		go worker(i, work, done)
	}

	for j := 0; j < 5; j++ {
		work <- j
	}
	close(work)

	for k := 0; k < 5; k++ {
		<-done
	}
	fmt.Println("All tasks procesed")
}
