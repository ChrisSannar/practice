package writeindex

// KeepPositives compacts s in place so the first k elements are the positive
// values (> 0) in their original order; it returns k. Elements past index k
// don't matter.
//
// START HERE — this is the pure write-index mechanic with the simplest possible
// keep-rule. The two pointers are NOT compared to each other: `read` visits every
// element and tests it on its own (`s[read] > 0`); `write` is just the cursor for
// where the next kept value goes, and only advances when you keep one.
func KeepPositives(s []int) int {
	write := 0
	for read := range s {
		if s[read] > 0 {
			s[write] = s[read]
			write++
		}
	}
	return write
}

// RemoveDuplicates takes an ascending-sorted slice and removes duplicates in
// place, so that the first k elements of s hold the unique values in order.
// It returns k (the count of unique values). Elements past index k don't matter.
//
// Use a slow "write" pointer trailing a fast "read" pointer: read scans every
// element; write only advances when you keep one. No extra slice, no map.
func RemoveDuplicates(s []int) int {
	if len(s) < 2 {
		return len(s)
	}
	write := 0
	for read := range s {
		if s[write] != s[read] {
			write++
			s[write] = s[read]
		}
	}
	return write + 1
}

// MoveZeroes moves all zeroes in s to the end while keeping the relative order
// of the non-zero elements. It mutates s in place and returns nothing.
//
// Same idiom: a write pointer marks the next slot for a non-zero value.
func MoveZeroes(s []int) {
	// Arrays with 1 item are already sorted
	if len(s) < 2 {
		return
	}

	// Anywhere we find a "Zero" we write over it with a non-zero element
	write := 0
	for read := range s {
		if s[read] != 0 {
			s[write] = s[read]
			write++
		}
	}

	// Once finished, we fill in the rest of the array
	for write < len(s) {
		s[write] = 0
		write++
	}
}

// RemoveElement removes every occurrence of val from s in place, so the first k
// elements of s are the survivors (order doesn't matter). It returns k. Elements
// past index k don't matter. s is not necessarily sorted.
//
// Same idiom again: read scans all, write keeps only the values that aren't val.
func RemoveElement(s []int, val int) int {
	if len(s) < 1 {
		return 0
	}
	write := 0
	for read := range s {
		if s[read] != val {
			s[write] = s[read]
			write++
		}
	}
	return write
}
