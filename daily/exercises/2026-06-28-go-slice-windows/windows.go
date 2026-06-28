package windows

// Take returns the first n elements of s.
// If n >= len(s), return all of s. If n <= 0, return an empty slice.
func Take(s []int, n int) []int {
	if len(s) <= n {
		return s
	}
	if n <= 0 {
		return []int{}
	}
	return s[0:n]
}

// Drop returns s without its first n elements.
// If n >= len(s), return an empty slice. If n <= 0, return all of s.
func Drop(s []int, n int) []int {
	if len(s) <= n {
		return []int{}
	}
	if n <= 0 {
		return s
	}
	return s[n:]
}

// Chunk splits s into consecutive groups of at most size elements.
// The last group may be shorter. size <= 0 returns an empty [][]int.
// An empty input returns an empty (non-nil) [][]int.
func Chunk(s []int, size int) [][]int {
	if len(s) == 0 || size <= 0 {
		return [][]int{}
	}

	result := [][]int{}
	idx := 0
	for idx+size >= len(s) {
		result = append(result, s[idx:idx+size])
		idx += size
	}
	result = append(result, s[idx:])

	return result
}
