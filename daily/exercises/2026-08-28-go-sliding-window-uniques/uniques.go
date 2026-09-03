package uniques

import (
	"fmt"
)

// FirstRepeatedIndex returns the index of the first element of a that has
// already appeared earlier in a (scanning left to right), or -1 if every
// element is distinct. Keep a set of values seen so far and check membership
// as you go.
//
// Example: FirstRepeatedIndex([]int{5, 1, 3, 1, 7}) -> 3
// Example: FirstRepeatedIndex([]int{1, 2, 3}) -> -1
func FirstRepeatedIndex(a []int) int {
	fmt.Println("FirstRepeatedIndex")

	set := make(map[int]bool)
	for idx, val := range a {
		alreadyAppeared := set[val]
		if alreadyAppeared {
			return idx
		} else {
			set[val] = true
		}
	}
	return -1
}

// LongestUniqueLen returns the length of the longest contiguous subarray of a
// whose elements are all distinct. Grow a window from the right into a set of
// its current values; when the incoming element is already in the set, shrink
// from the left (removing left elements from the set) until it isn't; track the
// largest window length seen.
//
// Example: LongestUniqueLen([]int{1, 2, 3, 1, 2}) -> 3
// Example: LongestUniqueLen([]int{7, 7, 7}) -> 1
// Example: LongestUniqueLen([]int{}) -> 0
func LongestUniqueLen(a []int) int {
	leftIdx := 0
	rightIdx := 0
	uniqueSet := make(map[int]bool)
	maxLen := 0

	for idx := range a {
		inSet := uniqueSet[a[idx]]
		if inSet {
			if rightIdx-leftIdx > maxLen {
				maxLen = rightIdx - leftIdx
			}
			delete(uniqueSet, a[idx])
			leftIdx = idx
		}
		uniqueSet[a[idx]] = true
		rightIdx++
	}

	if rightIdx-leftIdx == len(uniqueSet) {
		return len(uniqueSet)
	}
	return maxLen
}

// LengthOfLongestSubstring returns the length of the longest substring of s
// with no repeating characters (LeetCode 3). Same expand-right / shrink-left
// window as LongestUniqueLen, but over the runes of s.
//
// Example: LengthOfLongestSubstring("abcabcbb") -> 3
// Example: LengthOfLongestSubstring("bbbbb") -> 1
// Example: LengthOfLongestSubstring("pwwkew") -> 3
// Example: LengthOfLongestSubstring("") -> 0
func LengthOfLongestSubstring(s string) int {
	runes := []rune(s)
	leftIdx := 0
	rightIdx := 0
	uniqueSet := make(map[rune]bool)
	maxLen := 0

	for idx := range runes {
		inSet := uniqueSet[runes[idx]]
		if inSet {
			if rightIdx-leftIdx > maxLen {
				maxLen = rightIdx - leftIdx
			}
			delete(uniqueSet, runes[idx])
			leftIdx = idx
		}
		uniqueSet[runes[idx]] = true
		rightIdx++
	}

	if rightIdx-leftIdx == len(uniqueSet) {
		return len(uniqueSet)
	}
	return maxLen
}
