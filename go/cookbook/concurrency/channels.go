package main

import (
	"fmt"
	"time"
)

/*
* NOTE:
* - Always close from sender's side
* - Use `select` for multiple channels
* - Don't use as a mutex
* - Close your channel when done
* - Buffered Channels can exceed buffer consumption
 */
func main() {
	fmt.Println("Channels")
	// basic()
	// bufferedChannels()
	// closeChannel()
	selectChannel()
}

func basic() {
	ch := make(chan string)

	go func() {
		ch <- "Processing data form goroutine"
	}()

	msg := <-ch
	fmt.Println(msg)
}

func bufferedChannels() {
	ch := make(chan int, 2)

	ch <- 42
	ch <- 73

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func closeChannel() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch) // Only close channel if in goroutine
	}()

	for v := range ch {
		fmt.Println(v)
	}
}

func selectChannel() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "Data form channel 1"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "Data from channel 2"
	}()

	select {
	case msg1 := <-ch1:
		fmt.Println(msg1)
	case msg2 := <-ch2:
		fmt.Println(msg2)
	}
}
