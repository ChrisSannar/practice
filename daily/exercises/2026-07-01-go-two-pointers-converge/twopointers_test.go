package twopointers

import (
	"reflect"
	"testing"
)

func TestPairWithTarget(t *testing.T) {
	cases := []struct {
		name   string
		s      []int
		target int
		want   []int
	}{
		{"adjacent at front", []int{2, 7, 11, 15}, 9, []int{0, 1}},
		{"the two ends", []int{1, 5}, 6, []int{0, 1}},
		{"converge from both sides", []int{1, 3, 4, 5, 7, 11}, 9, []int{2, 3}},
		{"interior pair", []int{1, 2, 3, 4, 6}, 6, []int{1, 3}},
		{"no pair", []int{1, 2, 3}, 100, []int{}},
		{"no pair, near miss", []int{1, 2, 4, 8}, 7, []int{}},
		{"empty slice", []int{}, 5, []int{}},
		{"single element", []int{3}, 3, []int{}},
		{"negatives", []int{-5, -2, 0, 3, 6}, 1, []int{0, 4}},
	}
	for _, c := range cases {
		if got := PairWithTarget(c.s, c.target); !reflect.DeepEqual(got, c.want) {
			t.Errorf("%s: PairWithTarget(%v, %d) = %v, want %v", c.name, c.s, c.target, got, c.want)
		}
	}
}
