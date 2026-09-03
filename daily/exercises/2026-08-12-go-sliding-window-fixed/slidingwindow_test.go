package slidingwindow

import "testing"

func TestMaxSumK(t *testing.T) {
	cases := []struct {
		name string
		a    []int
		k    int
		want int
	}{
		{"mixed positives/negatives", []int{1, 9, -1, 3, 7}, 3, 11},
		{"whole middle is best", []int{2, 1, 5, 1, 3, 2}, 3, 9},
		{"all negatives true max is negative", []int{-4, -2, -3, -1}, 2, -4},
		{"k equals length", []int{1, 2, 3, 4}, 4, 10},
		{"k one picks the max element", []int{3, 1, 4, 1, 5}, 1, 5},
		{"tie for max does not matter for value", []int{1, 2, 1, 2, 1}, 2, 3},
		{"k zero invalid", []int{1, 2, 3}, 0, 0},
		{"k negative invalid", []int{1, 2, 3}, -1, 0},
		{"k exceeds length invalid", []int{1, 2, 3}, 5, 0},
	}
	for _, c := range cases {
		if got := MaxSumK(c.a, c.k); got != c.want {
			t.Errorf("%s: MaxSumK(%v, %d) = %d, want %d", c.name, c.a, c.k, got, c.want)
		}
	}
}

func TestMinSumK(t *testing.T) {
	cases := []struct {
		name string
		a    []int
		k    int
		want int
	}{
		{"mixed positives/negatives", []int{1, 9, -1, 3, 7}, 3, 9},
		{"all negatives true min is most negative", []int{-4, -2, -3, -1}, 2, -6},
		{"k equals length", []int{1, 2, 3, 4}, 4, 10},
		{"k one picks the min element", []int{3, 1, 4, 1, 5}, 1, 1},
		{"k zero invalid", []int{1, 2, 3}, 0, 0},
		{"k exceeds length invalid", []int{1, 2, 3}, 5, 0},
	}
	for _, c := range cases {
		if got := MinSumK(c.a, c.k); got != c.want {
			t.Errorf("%s: MinSumK(%v, %d) = %d, want %d", c.name, c.a, c.k, got, c.want)
		}
	}
}

func TestMaxSumKStart(t *testing.T) {
	cases := []struct {
		name string
		a    []int
		k    int
		want int
	}{
		{"max window starts at index 1", []int{1, 9, -1, 3, 7}, 3, 1},
		{"max window starts at index 3", []int{2, 1, 5, 1, 3, 2}, 3, 2},
		{"all negatives max is least negative window", []int{-4, -2, -3, -1}, 2, 2},
		{"k equals length only one window at 0", []int{1, 2, 3, 4}, 4, 0},
		{"k one max element earliest if tie", []int{3, 5, 4, 5, 2}, 1, 1},
		{"tie for max sum earliest index wins", []int{4, 1, 1, 4, 1}, 2, 0},
		{"k zero invalid", []int{1, 2, 3}, 0, -1},
		{"k negative invalid", []int{1, 2, 3}, -2, -1},
		{"k exceeds length invalid", []int{1, 2, 3}, 5, -1},
	}
	for _, c := range cases {
		if got := MaxSumKStart(c.a, c.k); got != c.want {
			t.Errorf("%s: MaxSumKStart(%v, %d) = %d, want %d", c.name, c.a, c.k, got, c.want)
		}
	}
}
