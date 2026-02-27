package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Worker Pools")
	basics()
}

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		fmt.Printf("Worker %d processing task %d\n", id, j)

		results <- j * 2
		fmt.Printf("Worker %d completed task %d\n", id, j)
	}
}

func basics() {
	const numWorkers = 3
	const numJobs = 5

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	var wg sync.WaitGroup

	for w := 0; w < numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	for j := 0; j < numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	wg.Wait()
	close(results)

	for result := range results {
		fmt.Println("Result:", result)
	}
}
