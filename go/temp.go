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
	// selectCases()
	/*
		arr := []int{1, 2, 3, 4, 5, 6}
		slc := arr[4:]
		fmt.Println(len(arr), len(slc))
		slc[0] = 0
		slc = append(slc, 7, 8, 9, 10, 11, 12, 13, 14)
		fmt.Println(arr, slc)
		fmt.Println(len(arr), len(slc))
	*/

	changeMap := func(merp map[string]int) {
		merp["a"] = 0
	}

	var mappers map[string]int = map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	changeMap(mappers)
	fmt.Println(mappers)

	changeArr := func(arr []int) {
		arr[0] = 0
	}
	passedArr := []int{1, 2, 3, 4}
	changeArr(passedArr[:])
	fmt.Println("Arr", passedArr)
	fmt.Println()

	var temp string = "yep"
	changeStr := func(str string) {
		str = "nope"
	}
	changeStr(temp)
	fmt.Println(temp)

	Printing()
}

type T struct {
	a int
	b float64
	c string
}

func Printing() {

	var x uint64 = 1<<64 - 1
	// `%d` prints the number, `%x` prints the hex
	fmt.Printf("%d %x; %d %x\n", x, x, int64(x), int64(x))

	// use `%v` as a "catch all" (value)
	t := &T{7, -2.35, "abc\tdef"}
	fmt.Printf("%v\n", t)
	fmt.Printf("%+v\n", t)
	fmt.Printf("%#v\n", t)
}

func Arrays() {
	Sum := func(a *[3]float64) (sum float64) {
		for _, v := range *a {
			sum += v
		}
		return
	}

	array := [...]float64{7.0, 8.5, 9.1}
	x := Sum(&array)
	fmt.Println("X", x)
	x2 := []int{1, 2, 3}
	fmt.Println("cap", cap(x2))
	x2 = append(x2, 4, 5)
	fmt.Println("cap2", cap(x2))
}
