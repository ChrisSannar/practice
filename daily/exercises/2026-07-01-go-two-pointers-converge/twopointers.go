package twopointers

import "fmt"

// PairWithTarget takes an ascending-sorted slice and a target. It returns the
// two indices {i, j} (with i < j) of a pair whose values sum to target. If no
// such pair exists, it returns an empty (non-nil) []int{}.
//
// Same for-l-<-r frame as before, but now move only ONE pointer each step,
// chosen by comparing the current sum to target. No nested loop, no map.
func PairWithTarget(s []int, target int) []int {
	fmt.Println("PairWithTarget test")
	for left, right := 0, len(s)-1; left < right; {
		sum := s[left] + s[right]
		if sum == target {
			return []int{left, right}
		}
		if sum < target {
			left++
		} else {
			right--
		}
	}
	return []int{}
}
