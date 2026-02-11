package main

import "fmt"

func main() {
	// 	var tempString string = "I am a temp string: racecar"
	// 	str := []byte(tempString)
	// 	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
	// 		str[i], str[j] = str[j], str[i]
	// 	}
	// 	fmt.Println(string(str))
	concurrencyPractice()
}

func basicChannel() {
	ch := make(chan string)
	go func() {
		ch <- "Hello from the goroutine"
	}()

	msg := <-ch
	fmt.Println(msg)
	fmt.Println("After message")
}

func basicRangeFeed() {
	jobs := make(chan int, 3)
	for i := 1; i <= 3; i++ {
		jobs <- i
	}
	close(jobs)

	for job := range jobs {
		fmt.Println("processing job", job)
	}
}

func multiChannel() {
	jobs2 := make(chan int)
	done := make(chan int)

	worker := func(id int) {
		for j := range jobs2 {
			fmt.Println("worker", id, "got job", j)
		}
		done <- id
	}

	for i := 1; i <= 2; i++ {
		go worker(i)
	}

	for j := 1; j <= 5; j++ {
		jobs2 <- j
	}
	close(jobs2)

	fmt.Println("Worker Done: ", <-done)
	fmt.Println("Worker Done: ", <-done)
}

func selectCases() {

	a := make(chan string)
	b := make(chan string)

	go func() { b <- "from B" }()
	go func() { a <- "from A" }()

	// We can use a select to determine what is coming in from the channel
	for i := 0; i < 2; i++ {
		select {
		case msg := <-a:
			fmt.Println(msg)
		case msg := <-b:
			fmt.Println(msg)
		}
	}
}

func mergeChannels() {

}

func concurrencyPractice() {
	fmt.Println("concurrencyPractice")

	// 1. Basic channel example
	// basicChannel()

	// 2. A feed and range of values into a channel
	// basicRangeFeed()

	// 3. Running work form multiple channels
	// multiChannel()

	// 4. Can select from current channel output
	selectCases()
}
