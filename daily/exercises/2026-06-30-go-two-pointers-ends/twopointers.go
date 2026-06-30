package twopointers

// import (
// 	"fmt"
// )

// IsPalindrome reports whether s reads the same forward and backward.
// Use two pointers: one at each end, walking toward the middle, comparing as
// you go. An empty slice and a single element are palindromes.
func IsPalindrome(s []int) bool {
	for leftIdx, rightIdx := 0, len(s)-1; leftIdx < rightIdx; {
		if s[leftIdx] != s[rightIdx] {
			return false
		}

		leftIdx++
		rightIdx--
	}
	return true
}

// ReverseInPlace reverses s in place (it mutates the caller's slice and returns
// nothing). Use two pointers: one at each end, swap, then move both inward.
func ReverseInPlace(s []int) {
	for leftIdx, rightIdx := 0, len(s)-1; leftIdx < rightIdx; {
		s[leftIdx], s[rightIdx] = s[rightIdx], s[leftIdx]

		leftIdx++
		rightIdx--
	}
}
