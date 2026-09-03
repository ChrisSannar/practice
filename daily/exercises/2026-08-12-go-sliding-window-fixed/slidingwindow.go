package slidingwindow

import (
	"math"
)

func sum(a []int) int {
	result := 0
	for _, val := range a {
		result += val
	}
	return result
}

// MaxSumK returns the maximum sum of any contiguous subarray of a of length k.
// If k <= 0 or k > len(a), returns 0.
//
// Example: MaxSumK([]int{1, 9, -1, 3, 7}, 3) -> 11
func MaxSumK(a []int, k int) int {
	if k <= 0 || k > len(a) {
		return 0
	}

	result := math.MaxInt * -1
	for i := 0; i+k < len(a)+1; i++ {
		s := sum(a[i : i+k])
		if s > result {
			result = s
		}
	}
	return result
}

// MinSumK returns the minimum sum of any contiguous subarray of a of length k.
// If k <= 0 or k > len(a), returns 0.
//
// Example: MinSumK([]int{1, 9, -1, 3, 7}, 3) -> 9
func MinSumK(a []int, k int) int {
	if k <= 0 || k > len(a) {
		return 0
	}

	result := math.MaxInt
	for i := 0; i+k < len(a)+1; i++ {
		s := sum(a[i : i+k])
		if s < result {
			result = s
		}
	}
	return result
}

// MaxSumKStart returns the starting index of the window of length k whose sum
// is maximum. If multiple windows tie for the max, returns the earliest index.
// If k <= 0 or k > len(a), returns -1.
//
// Example: MaxSumKStart([]int{1, 9, -1, 3, 7}, 3) -> 1
func MaxSumKStart(a []int, k int) int {
	if k <= 0 || k > len(a) {
		return -1
	}

	maxVal := math.MaxInt * -1
	idx := 0
	for i := 0; i+k < len(a)+1; i++ {
		s := sum(a[i : i+k])
		if s > maxVal {
			idx = i
			maxVal = s
		}
	}
	return idx
}
