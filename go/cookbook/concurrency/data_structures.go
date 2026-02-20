package main

import (
	"fmt"
	"sync"
)

/*
* NOTE:
* - Mutex is for critical sections for safe read/writes
* - `sync.RWMutex` is good for multi read structures
* - Map is easier to manage, but less performant
* - Don't forget to unlock if your return early
* - Don't use `sync.Map` indiscriminately
* - Don't assume map is always safe
* - Use fine-grained for highly parallelized apps
* - Profile!
* - Use `sync.Pool` to reduce allocation overhead
 */
func main() {
	fmt.Println("Data Structures")
	// mutexMaps()
	// rwMutex()
	syncMap()
}

func mutexMaps() {
	var mu sync.Mutex
	m := make(map[string]int)

	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1) // "We're adding to the wait group!"
		go func(i int) {
			defer wg.Done() // We're finished after this function is
			mu.Lock()       // Ok, now don't touch!
			m[fmt.Sprintf("task%d", i)] = i
			mu.Unlock() // Ok, we're done now, someone else can use it
		}(i)
	}
	wg.Wait() // Once all goroutines are finished...

	mu.Lock()
	fmt.Println("Data contents:", m) // We can safely print out contents!
	mu.Unlock()
}

func rwMutex() {
	var rwMu sync.RWMutex
	m := make(map[string]int)

	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			rwMu.Lock()
			m[fmt.Sprintf("task%d", i)] = i
			rwMu.Unlock()
		}(i)
	}

	// This next block is a race condition
	// wg.Wait() // <- but adding this fixes that race condition
	wg.Add(1)
	go func() {
		defer wg.Done()
		rwMu.RLock() // Only lock for reading
		fmt.Println("Reading data contents:", m)
		rwMu.RUnlock()
	}()

	wg.Wait()
}

func syncMap() {
	var sm sync.Map

	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			sm.Store(fmt.Sprintf("task%d", i), i)
		}(i)
	}

	wg.Wait()

	sm.Range(func(key, value interface{}) bool {
		fmt.Printf("%s: %d\n", key, value)
		return true
	})
}
