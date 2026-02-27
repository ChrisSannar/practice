package main

import (
	"fmt"
	"sync"
)

/**
*	NOTE:
*	- Defer `mu.Unlock()` to make sure we don't deadlock
*	- `RWMutex` is better for high read-to-write ratios
*	- Keep critical sections (between `Lock` and `Unlock`) short
*	- Overuse affects performance
*	- Don't copy a mutex after use (will panic)
*	- Don't lock for too long
*	- Use channels for communication
*	- Benchmark to decide between `Mutex` and `RWMutex`
 */
func main() {
	fmt.Println("Mutex")
	basics()
	readWriteMutex()
}

func basics() {
	var mu sync.Mutex
	value := 0

	increment := func(wg *sync.WaitGroup, mu *sync.Mutex) {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			mu.Lock()
			value++
			mu.Unlock()
		}
	}

	var wg sync.WaitGroup

	wg.Add(2)
	go increment(&wg, &mu)
	go increment(&wg, &mu)

	wg.Wait()
	fmt.Printf("Final value: %d\n", value)
}

func readWriteMutex() {
	var mu sync.RWMutex
	value := 0

	read := func(wg *sync.WaitGroup, mu *sync.RWMutex) {
		defer wg.Done()
		mu.RLock()
		defer mu.RUnlock()
		fmt.Println("Reading value:", value)
	}

	write := func(wg *sync.WaitGroup, mu *sync.RWMutex) {
		defer wg.Done()
		mu.Lock()
		value++
		mu.Unlock()
	}

	var wg sync.WaitGroup

	wg.Add(3)
	go write(&wg, &mu)
	go read(&wg, &mu)
	go read(&wg, &mu)

	wg.Wait()
}
