package main

import (
	"fmt"
	"time"
)

/*
* NOTE:
* - Select is best for multi channel
* - Use `default` to keep from blocking
* - `time.After` creates a channel so careful of the overhead
* - `select` will block if there are no defauls
* - Handle channel closures correctly
* - Use `select` over polling
* - Use buffered channels to preved blocking
 */
func main() {
	fmt.Println("Select")
	// basics()
	// defaultCase()
	selectTimeout()
}

func basics() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "Processing task 1"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "Processing task 2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("Recieved:", msg1)
		case msg2 := <-ch2:
			fmt.Println("Recieved:", msg2)
		}
	}
}

func defaultCase() {
	ch := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- "Data Processing complete"
	}()

	for {
		select {
		case msg := <-ch:
			fmt.Println("Recieved:", msg)
			return
		default:
			fmt.Println("Waiting for data...")
			time.Sleep(500 * time.Millisecond)
		}

	}
}

func selectTimeout() {
	ch := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- "Data processing complete"
	}()

	select {
	case msg := <-ch:
		fmt.Println("Recieved:", msg)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout: no data recieved")
	}
}
