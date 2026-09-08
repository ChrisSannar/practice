package shrink

// ShrinkPastDuplicate returns the new left edge of the window a[left:right]
// after removing a[left], then a[left+1], and so on — one index at a time —
// stopping the instant the element equal to a[right] has been removed. The
// window a[left:right] (indices left through right-1) starts with no
// duplicates in it. If nothing in a[left:right] equals a[right], left is
// returned unchanged.
//
// Example: ShrinkPastDuplicate([]int{1, 2, 5, 9, 5}, 0, 4) -> 3
// Example: ShrinkPastDuplicate([]int{5, 1, 2, 9, 5}, 0, 4) -> 1
// Example: ShrinkPastDuplicate([]int{1, 2, 3, 4, 5}, 0, 4) -> 0
func ShrinkPastDuplicate(a []int, left, right int) int {
	// TODO: implement
	for idx := left; idx < right; idx++ {
		if a[idx] == a[right] {
			left++
		}
	}
	return left
}

// LongestUniqueLen returns the length of the longest contiguous subarray of a
// whose elements are all distinct. Grow a window from the right into a set of
// its current values; on a duplicate, shrink from the left one element at a
// time (ShrinkPastDuplicate's mechanic) until it clears, then add the
// incoming element; track the largest window length seen.
//
// Example: LongestUniqueLen([]int{0, 1, 2, 0, 3}) -> 4
// Example: LongestUniqueLen([]int{7, 7, 7}) -> 1
// Example: LongestUniqueLen([]int{}) -> 0
func LongestUniqueLen(a []int) int {
	// TODO: implement
	return 0
}

// LengthOfLongestSubstring returns the length of the longest substring of s
// with no repeating characters (LeetCode 3). Same expand-right / shrink-left
// window as LongestUniqueLen, but over the runes of s.
//
// Example: LengthOfLongestSubstring("dvdf") -> 3
// Example: LengthOfLongestSubstring("bbbbb") -> 1
// Example: LengthOfLongestSubstring("") -> 0
func LengthOfLongestSubstring(s string) int {
	// TODO: implement
	return 0
}
