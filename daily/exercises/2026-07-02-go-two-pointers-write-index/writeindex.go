package writeindex

// RemoveDuplicates takes an ascending-sorted slice and removes duplicates in
// place, so that the first k elements of s hold the unique values in order.
// It returns k (the count of unique values). Elements past index k don't matter.
//
// Use a slow "write" pointer trailing a fast "read" pointer: read scans every
// element; write only advances when you keep one. No extra slice, no map.
func RemoveDuplicates(s []int) int {
	// TODO: slow write pointer, fast read pointer.
	return 0
}

// MoveZeroes moves all zeroes in s to the end while keeping the relative order
// of the non-zero elements. It mutates s in place and returns nothing.
//
// Same idiom: a write pointer marks the next slot for a non-zero value.
func MoveZeroes(s []int) {
	// TODO: slow write pointer, fast read pointer.
}

// RemoveElement removes every occurrence of val from s in place, so the first k
// elements of s are the survivors (order doesn't matter). It returns k. Elements
// past index k don't matter. s is not necessarily sorted.
//
// Same idiom again: read scans all, write keeps only the values that aren't val.
func RemoveElement(s []int, val int) int {
	// TODO: slow write pointer, fast read pointer.
	return 0
}
