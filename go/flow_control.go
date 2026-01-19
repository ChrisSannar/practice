package main

import (
	"fmt"
	"math"
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

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v:= math.Pow(x, n); v < lim {
		return v
	}
	return lim
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
}
