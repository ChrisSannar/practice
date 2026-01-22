package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func normalLoop() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += 1
	}
	fmt.Println(sum)
}

func whileLoop() {
	sum := 1
	for sum < 1000 {
		fmt.Println(sum)
		sum += sum
	}
	fmt.Println("done", sum)
}

func foreverLoop() {
	for {}
}

// An Example of conditionals as well as a hint of recursion
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

// The 'v' variable is only available inside of the 'if' statements.
func pow(x, n, lim float64) float64 {
	if v:= math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// it cant be accessed here
	return lim
}

func MySqrt(x float64) float64 {
	z := float64(1)
	i := 0
	for i < 10 {
		i += 1
		z -= (z*z - x) / (2 * z)
	}
	return z
}

// Switches automatically break when encountering a case
func switches() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("MacOS.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s. \n", os)
	}
}

// Switches stop on the first case that it encounters
func switches2() {
	fmt.Println("When is Saturday?")
	today := time.Now().Weekday()
	// Curious... Why does additition work here?
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today")
	case today + 1:
		fmt.Println("Tomorrow")
	case today + 2:
		fmt.Println("In two days")
	default:
		fmt.Printf("It's %s? Too far away\n", today)
	}
}

// Evaluations can be made as well. Anything that evaluates to "true" will trigger the case
func switches3() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good Morning!")
	case t.Hour() > 17:
		fmt.Println("Good Afternoon!")
	default: 
		fmt.Println("Good Evening")
	}
}

// "defer" will evaluate the value first, but then execute when the function is finished
func deferals() {
	x := 2
	defer fmt.Println("This is evaluated, then printed", x)
	x += 2

	fmt.Println("This is called first", x)
}

// "defer" executes in a stack so it follows a "LIFO" order (Last In, First Out)
func deferals2() {
	fmt.Println("Counting...")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

func main() {
	normalLoop()
	whileLoop()
	// foreverLoop()
	fmt.Println(sqrt(4), sqrt(-9))
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
	fmt.Println(MySqrt(2))
	switches()
	switches2()
	switches3()
	deferals()
	deferals2()
}
