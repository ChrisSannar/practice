package slidingwindow

import (
	"fmt"
	"math"
)

func PrintArr(a []int) {
	if len(a) > 10 {
		return
	}
	fmt.Println(a)
}

// WindowSums returns the sums of every contiguous subarray of a of length k,
// in left-to-right order. The result has len(a)-k+1 entries. If k <= 0 or
// k > len(a), returns nil.
//
// Example: WindowSums([]int{1, 2, 3, 4, 5}, 3) -> [6, 9, 12]
func WindowSums(a []int, k int) []int {
	if k > len(a) || k <= 0 {
		return nil
	}
	sums := make([]int, 0)
	sum := 0
	for i := range k {
		sum += a[i]
	}
	sums = append(sums, sum)
	for i := k; i < len(a); i++ {
		sum += a[i] - a[i-k]
		sums = append(sums, sum)
	}

	return sums
}

// MaxSumKAndStart returns the maximum sum of any contiguous subarray of a of
// length k and the starting index of an earliest window achieving it. If
// multiple windows tie for the max, the earliest start wins. If k <= 0 or
// k > len(a), returns (0, -1).
//
// Example: MaxSumKAndStart([]int{2, 1, 5, 1, 3, 2}, 3) -> (9, 2)
func MaxSumKAndStart(a []int, k int) (int, int) {
	if k > len(a) || k <= 0 {
		return 0, -1
	}
	runningSum := 0
	for i := range k {
		runningSum += a[i]
	}
	maxSum := runningSum
	maxSumIdx := 0
	for i := k; i < len(a); i++ {
		runningSum += a[i] - a[i-k]
		if runningSum > maxSum {
			maxSum = runningSum
			maxSumIdx = i - k + 1
		}
	}

	return maxSum, maxSumIdx
}

// MinLenSumAtLeast returns the length of the shortest contiguous subarray of a
// whose sum is >= target, or 0 if no such subarray exists. a contains only
// positive integers, so growing always raises the sum and shrinking always
// lowers it. target <= 0 yields 0.
//
// Example: MinLenSumAtLeast([]int{2, 3, 1, 2, 4, 3}, 7) -> 2
func MinLenSumAtLeast(a []int, target int) int {
	if target <= 0 {
		return 0
	}

	runningSum := 0
	leftIdx := 0
	minLen := math.MaxInt
	for rightIdx, val := range a {
		runningSum += val
		for runningSum >= target {
			if rightIdx-leftIdx < minLen {
				minLen = rightIdx - leftIdx
			}
			runningSum -= a[leftIdx]
			leftIdx++
		}
	}

	if minLen == math.MaxInt {
		return 0
	}
	return minLen + 1
}
