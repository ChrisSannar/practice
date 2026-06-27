package slicesmaps

import (
	"sort"
)

// Chunk splits s into consecutive groups of at most size elements.
// See SPEC.md for the size <= 0 and empty-input contracts.
func Chunk(s []int, size int) [][]int {
	//
	// result := make([][]int, size) // Our result is going to be our "Chunk" size, because that's how many "Chunks" we want
	// if len(s) == 0 || size == 0 {
	// 	return result
	// }
	//
	// chunkSize := len(s) / size
	// numberOfChunks := len(s) / chunkSize
	//
	// fmt.Println("result:", result)
	// fmt.Println("chunkSize:", chunkSize)
	// fmt.Println("numberOfChunks:", numberOfChunks)
	//
	// // startIdx := 0
	// for idx := range chunkSize / len(s) {
	// 	fmt.Println(idx)
	// }
	//
	// return result
	return nil
}

// Unique returns the elements of s with duplicates removed, keeping the order
// in which each value first appears.
func Unique(s []int) []int {
	result := make([]int, 0)
	mappers := make(map[int]bool)
	for _, val := range s {
		inMap := mappers[val]
		if !inMap {
			result = append(result, val)
			mappers[val] = true
		}
	}
	return result
}

// Intersection returns the values that appear in both a and b, with no
// duplicates, in the order they first appear in a.
func Intersection(a, b []int) []int {
	// Optimize size for runtime. Could we use pointers to keep from initialization?
	var largerArr []int
	var smallerArr []int
	if len(a) > len(b) {
		largerArr, smallerArr = a, b
	} else {
		largerArr, smallerArr = b, a
	}
	result := make([]int, 0)

	// using `mappers` to keep a run time of O(n). Larger space, but that's usually a better trade off
	mappers := make(map[int]bool)
	for _, val := range largerArr {
		mappers[val] = true
	}
	for _, val := range smallerArr {
		if mappers[val] {
			result = append(result, val)
		}
	}

	// Should have read the instructions, but I was thinking too far ahead
	sort.Ints(result)
	return result
}

// Merge combines any number of maps into one, summing the values of shared keys.
func Merge(maps ...map[string]int) map[string]int {
	THE_MAP := make(map[string]int)

	// This one has O(n^2) complexity. Don't know how to fix it otherwise
	for _, mappers := range maps {
		for key, val := range mappers {
			THE_MAP[key] = THE_MAP[key] + val
		}
	}

	return THE_MAP
}
