package main

import (
	"fmt"
	"time"
)

// Returns the index in `s` of item `x`
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		// v and x are of type `T` which is of the type `comparable`
		if v == x { // Because it's `comparable`, this works
			return i
		}
	}
	return -1
}

type List[T any] struct {
	next *List[T]
	val  T
}

func buildList[T any](items []T) *List[T] {
	if len(items) == 0 {
		return nil
	}

	var head *List[T] = &List[T]{val: items[0], next: nil}
	var current *List[T] = head

	for idx := 1; idx < len(items); idx++ {
		current.next = &List[T]{val: items[idx]}
		current = current.next
	}

	return head
}

func printList[T any](node List[T]) {
	fmt.Print(node.val)
	if node.next != nil {
		fmt.Print(" -> ")
		printList(*node.next)
	}
}

func generics() {
	fmt.Println("Generics")
	si := []int{10, 20, 15, -10}
	fmt.Println(Index(si, 15))

	ss := []string{"foo", "bar", "baz"}
	fmt.Println(Index(ss, "hello"))
	fmt.Println()

	var l List[string] = *buildList([]string{"a", "b", "c"})
	fmt.Println(l)
	fmt.Println(l.next)
	fmt.Println(l.next.next)
	fmt.Println(l.next.next.next)
	printList(l)
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func sum(s []int, c chan int) {
	sum := 0
	// Add up all the numbers in your array
	for _, v := range s {
		sum += v
	}
	c <- sum // Channel it out
}

func fibonacci2(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c) // ALWAYS CLOSE THE CHANNEL
	// Else you'll get a panic
}

func fib(c, quit chan int) {
	x, y := 0, 1
	for {
		// `select` "pauses" until a routine is ready
		select {
		// our sequence continuse, since it's always ready
		case c <- x:
			x, y = y, x+y
			// `quit` only happens when it's ready
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func selectDefault() {
	start := time.Now()
	tick := time.Tick(100 * time.Millisecond)  // Channels every 100ms
	boom := time.After(500 * time.Millisecond) // Channels after 500ms
	// Returns the how long since `start`
	elapsed := func() time.Duration {
		return time.Since(start).Round(time.Millisecond)
	}
	for {
		select {
		case <-tick:
			fmt.Printf("[%6s] tick.\n", elapsed())
		case <-boom:
			fmt.Printf("[%6s] BOOM!\n", elapsed())
			return
		// This will run regardless unless there is a break
		default:
			fmt.Printf("[%6s].\n", elapsed())
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func concurrency() {
	fmt.Println("Concurrency: ")
	// This creates a "goroutine" that has access to the same memory pool
	// go say("world")
	// say("hello")

	// Channels work similar to pipes, but we don't have to worry about locking
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)     // channels must be "made"
	go sum(s[:len(s)/2], c) // Take the first half...
	go sum(s[len(s)/2:], c) // And the second half...
	x, y := <-c, <-c        // The channels are delivered in the order they were created

	fmt.Println(x, y, x+y)
	fmt.Println()

	// Channels can have a buffered size
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	// ch <- 3 // Overfills the buffer
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	// fmt.Println(<-ch)	// Also panics since there's nothing to pull

	c2 := make(chan int, 98)   // Keep in mind, fib(98) will overflow `int` and give you negative numbers
	go fibonacci2(cap(c2), c2) // `cap` is just capacity: 98
	for i := range c2 {        // range can be used on channels till it `close`s
		// fmt.Println(i)
		i *= 2
	}
	fmt.Println()

	c3 := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c3)
		}
		// In this case, quite only happens after all the values have been printed
		quit <- 0
	}()
	fib(c3, quit)
	fmt.Println()

	selectDefault()
}

func main() {
	// generics()
	concurrency()
}
