package basics

import (
	"fmt"
	"strings"
)

// Max returns the larger of a and b. If equal, returns either.
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Sum returns the sum of all ints in nums. The sum of an empty slice is 0.
func Sum(nums []int) int {
	total := 0
	for _, val := range nums {
		total += val
	}
	return total
}

// Reverse returns s with its characters in reverse order.
// Must handle multi-byte (Unicode) characters correctly, not just bytes.
func Reverse(s string) string {
	result := ""
	runes := []rune(s)
	for i := len(runes) - 1; i >= 0; i-- {
		result += string(runes[i])
	}
	return result
}

// CountVowels returns how many vowels (a, e, i, o, u) appear in s,
// case-insensitively.
func CountVowels(s string) int {
	vowels := "aeiouAEIOU"
	count := 0
	for _, val := range s {
		if strings.Contains(vowels, string(val)) {
			count++
		}
	}
	return count
}

// WordCount returns a map from each whitespace-separated word in s to the
// number of times it appears. An empty/blank string returns an empty map.
func WordCount(s string) map[string]int {
	// TODO
	f := strings.Fields((s))
	var mappers map[string]int = make(map[string]int)
	for _, val := range f {
		_, mapVal := mappers[val]
		if mapVal {
			mappers[val]++
		} else {
			mappers[val] = 1
		}
	}
	return mappers
}

// FizzBuzz returns a slice of length n (for n >= 0) where index i (1-based,
// so the first element corresponds to 1) is:
//   - "FizzBuzz" if the number is divisible by both 3 and 5
//   - "Fizz" if divisible by 3
//   - "Buzz" if divisible by 5
//   - the number as a string otherwise
func FizzBuzz(n int) []string {
	result := []string{}

	for i := 1; i <= n; i++ {
		div3 := i%3 == 0
		div5 := i%5 == 0
		if div3 && div5 {
			result = append(result, "FizzBuzz")
			continue
		}
		if div3 {
			result = append(result, "Fizz")
			continue
		}
		if div5 {
			result = append(result, "Buzz")
			continue
		}
		result = append(result, fmt.Sprintf("%d", i))
	}
	return result
}
