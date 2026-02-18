package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
* NOTE:
* - Use Atomic operations for lock-free concurrency control
* - Use locks if atomic types don't fit
* - Opt for `atomic.Value` for complex types
* - Avoid using atomics for sizes larger than 64 bits
* - Limit use for performance
* - Ensure shared all shared data operations are atomic
* - Atomic > mutex/semephores
 */
func main() {
	fmt.Println("Atomic")
	// basics()
	// compareAndSwap()
	atomicValue()
}

func basics() {
	var value int32
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt32(&value, 1)
		}()
	}

	wg.Wait() // Ensures the group is done with all async operations
	fmt.Printf("Final value: %d\n", value)
}

func compareAndSwap() {
	var value int32 = 42

	swapped := atomic.CompareAndSwapInt32(&value, 42, 73)
	fmt.Printf("Swapped: %v, Value: %d\n", swapped, value)
}

type User struct {
	Name string
	Age  int
}

func atomicValue() {
	var value atomic.Value

	// Store initial user.
	user1 := User{Name: "Alex", Age: 37}
	value.Store(user1)

	// Load and print.
	current := value.Load().(User)
	fmt.Printf("Current user: %+v\n", current)

	// Update with new user data.
	user2 := User{Name: "Vera", Age: 36}
	value.Store(user2)

	// Load and print updated data.
	updated := value.Load().(User)
	fmt.Printf("Updated user: %+v\n", updated)
}
