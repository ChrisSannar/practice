package main

import (
	"fmt"
	"sync"
	"time"
)

/*
* NOTE:
* - Always call `Release` after `Aquire` to avoid resource leaks
* - Use buffered channels
* - Be aware of deadlocks
* - Don't close channels with Sempahores
* - Choose a minimal buffer size
 */
func main() {
	fmt.Println("Semaphore")
	basics()
}

type Semaphore struct {
	Channel chan struct{}
}

func (s *Semaphore) Aquire() {
	s.Channel <- struct{}{}
}

func (s *Semaphore) Release() {
	<-s.Channel
}

func basics() {
	sem := &Semaphore{Channel: make(chan struct{}, 3)}
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			sem.Aquire()
			defer sem.Release()

			fmt.Printf("Processing task %d\n", id)
			time.Sleep(1 * time.Second)
		}(i)
	}

	wg.Wait()
}
