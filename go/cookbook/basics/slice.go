package main

import (
	"fmt"
	"slices"
)

/*
* NOTE:
* - Pre-allocate slices using make (to avoid multiples)
* - If large, be careful with `append` (use `cap()`)
* - Slices are passed by reference (header includes: pointer, len, cap)
* */
func main() {
	fmt.Println("Go Slices")
	// basic()
	// sliced()
	// appendSlice()
	// copySlice()
	// modSlice()
	utilSlice()
}

func basic() {
	slice := make([]int, 3, 10)
	fmt.Println(slice)
}

func sliced() {
	arr := [5]int{1, 2, 3, 4, 5}
	slice := arr[1:4]
	fmt.Println(slice)
}

func appendSlice() {
	slice := []int{1, 2, 3}
	slice = append(slice, 4)
	slice = append(slice, 5, 6, 7)
	fmt.Println(slice)
}

func copySlice() {
	source := []int{1, 2, 3, 4}
	destination := make([]int, len(source))
	copy(destination, source)
	fmt.Println(destination)
}

func modSlice() {
	slice := []int{1, 2, 3, 4, 5}
	for i := range slice {
		slice[i] *= 2
	}
	fmt.Println(slice)
}

func utilSlice() {
	numbers := []int{3, 1, 4, 1, 5, 9, 2, 6}

	fmt.Printf("Contains 5: %t\n", slices.Contains(numbers, 5))
	fmt.Printf("Index of 4: %d\n", slices.Index(numbers, 4))

	sortedNumbers := slices.Clone(numbers)
	slices.Sort(sortedNumbers)
	fmt.Printf("Sorted: %v\n", sortedNumbers)
	binarySearchRes, _ := slices.BinarySearch(sortedNumbers, 5)
	fmt.Printf("Binary search for 5: %d\n", binarySearchRes)

	inserted := slices.Insert(numbers, 2, 99, 100)
	fmt.Printf("AFter insert: %v\n", inserted)
	deleted := slices.Delete(numbers, 1, 3)
	fmt.Printf("After delete: %v\n", deleted)
	other := []int{3, 1, 4, 1, 5, 9, 2, 6}
	fmt.Printf("Equal to other: %t\n", slices.Equal(numbers, other))
	reversed := slices.Clone(numbers)
	slices.Reverse(reversed)
	fmt.Printf("Reversed: %v\n", reversed)
}
