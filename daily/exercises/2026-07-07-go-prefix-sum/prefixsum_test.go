package prefixsum

import (
	"reflect"
	"testing"
)

func TestBuildPrefixSum(t *testing.T) {
	cases := []struct {
		name string
		a    []int
		want []int
	}{
		{"ordinary", []int{1, 2, 3, 4}, []int{0, 1, 3, 6, 10}},
		{"single", []int{5}, []int{0, 5}},
		{"empty", []int{}, []int{0}},
		{"negatives", []int{-1, 2, -3}, []int{0, -1, 1, -2}},
	}
	for _, c := range cases {
		got := BuildPrefixSum(c.a)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("%s: BuildPrefixSum(%v) = %v, want %v", c.name, c.a, got, c.want)
		}
	}
}

func TestRangeSum(t *testing.T) {
	pre := []int{0, 1, 3, 6, 10, 15} // from []int{1, 2, 3, 4, 5}
	cases := []struct {
		name string
		i, j int
		want int
	}{
		{"first three", 0, 2, 6},
		{"middle", 1, 3, 9},
		{"single element", 2, 2, 3},
		{"whole array", 0, 4, 15},
		{"last element", 4, 4, 5},
	}
	for _, c := range cases {
		if got := RangeSum(pre, c.i, c.j); got != c.want {
			t.Errorf("%s: RangeSum(pre, %d, %d) = %d, want %d", c.name, c.i, c.j, got, c.want)
		}
	}
}

func TestHasSubarraySum(t *testing.T) {
	cases := []struct {
		name   string
		a      []int
		target int
		want   bool
	}{
		{"pair in middle", []int{1, 2, 3}, 5, true},
		{"whole array", []int{1, 2, 3}, 6, true},
		{"no match", []int{1, 2, 3}, 100, false},
		{"crosses negative", []int{-1, -1, 1}, 0, true},
		{"single zero", []int{0}, 0, true},
		{"empty", []int{}, 0, false},
	}
	for _, c := range cases {
		if got := HasSubarraySum(c.a, c.target); got != c.want {
			t.Errorf("%s: HasSubarraySum(%v, %d) = %v, want %v", c.name, c.a, c.target, got, c.want)
		}
	}
}

func TestCountSubarraysWithSum(t *testing.T) {
	cases := []struct {
		name   string
		a      []int
		target int
		want   int
	}{
		{"two matches", []int{1, 1, 1}, 2, 2},
		{"two different subarrays", []int{1, 2, 3}, 3, 2},
		{"no match", []int{1}, 0, 0},
		{"empty", []int{}, 0, 0},
		{"crosses negative", []int{-1, -1, 1}, 0, 1},
	}
	for _, c := range cases {
		if got := CountSubarraysWithSum(c.a, c.target); got != c.want {
			t.Errorf("%s: CountSubarraysWithSum(%v, %d) = %d, want %d", c.name, c.a, c.target, got, c.want)
		}
	}
}
