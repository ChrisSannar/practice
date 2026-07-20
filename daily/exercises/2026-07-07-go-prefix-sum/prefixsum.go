package prefixsum

import "fmt"

// BuildPrefixSum returns pre where pre[0] = 0 and pre[i] = pre[i-1] + a[i-1]
// for i = 1..len(a). len(pre) == len(a)+1.
//
// Example: BuildPrefixSum([]int{1, 2, 3, 4}) -> []int{0, 1, 3, 6, 10}
func BuildPrefixSum(a []int) []int {
	result := []int{0}
	for idx, val := range a {
		result = append(result, result[idx]+val)
	}
	return result
}

// RangeSum returns the sum of a[i..j] inclusive, in O(1), using a prefix
// slice built by BuildPrefixSum.
//
// Example: pre = []int{0, 1, 3, 6, 10} (from a = []int{1, 2, 3, 4});
// RangeSum(pre, 1, 3) -> 9 (i.e. a[1]+a[2]+a[3] = 2+3+4)
func RangeSum(pre []int, i, j int) int {
	return pre[j+1] - pre[i]
}

// HasSubarraySum reports whether any contiguous subarray of a sums to target.
//
// Example: HasSubarraySum([]int{1, 2, 3}, 5) -> true (2+3 = 5)
func HasSubarraySum(a []int, target int) bool {
	// prefixArray := BuildPrefixSum(a)
	fmt.Println("HasSubarraySum", a)

	// TODO
	return false
}

// CountSubarraysWithSum returns the number of contiguous subarrays of a that
// sum to target.
//
// Example: CountSubarraysWithSum([]int{1, 2, 3}, 3) -> 2 ([1,2] and [3])
func CountSubarraysWithSum(a []int, target int) int {
	fmt.Println("CountSubarraysWithSum")
	// TODO
	return 0
}
