package slidingwindow

import "testing"

func TestWindowSums(t *testing.T) {
	cases := []struct {
		name string
		a    []int
		k    int
		want []int
	}{
		{"ordinary mixed", []int{1, 2, 3, 4, 5}, 3, []int{6, 9, 12}},
		{"single element windows k one", []int{3, 1, 4, 1, 5}, 1, []int{3, 1, 4, 1, 5}},
		{"k equals length one window", []int{1, 2, 3, 4}, 4, []int{10}},
		{"all negatives", []int{-4, -2, -3, -1}, 2, []int{-6, -5, -4}},
		{"k zero invalid", []int{1, 2, 3}, 0, nil},
		{"k negative invalid", []int{1, 2, 3}, -1, nil},
		{"k exceeds length invalid", []int{1, 2, 3}, 5, nil},
		{"empty array any k invalid", []int{}, 1, nil},
	}
	for _, c := range cases {
		got := WindowSums(c.a, c.k)
		if !intsEqual(got, c.want) {
			t.Errorf("%s: WindowSums(%v, %d) = %v, want %v", c.name, c.a, c.k, got, c.want)
		}
	}
}

// A correctness test large enough that an O(n*k) recompute-from-scratch loop
// takes several seconds while an O(n) seed+slide finishes in microseconds.
// If this test visibly hangs, the recompute habit from 4a is back: use the slide.
func TestWindowSumsLarge(t *testing.T) {
	const n, k = 100000, 50000
	a := make([]int, n)
	for i := range a {
		a[i] = i
	}
	got := WindowSums(a, k)
	if len(got) != n-k+1 {
		t.Fatalf("len = %d, want %d", len(got), n-k+1)
	}
	if got[0] != 1249975000 {
		t.Errorf("first window sum = %d, want 1249975000", got[0])
	}
	if got[n-k] != 3749975000 {
		t.Errorf("last window sum = %d, want 3749975000", got[n-k])
	}
}

func TestMaxSumKAndStart(t *testing.T) {
	cases := []struct {
		name      string
		a         []int
		k         int
		wantSum   int
		wantStart int
	}{
		{"max window starts at index 2", []int{2, 1, 5, 1, 3, 2}, 3, 9, 2},
		{"max window starts at index 1", []int{1, 9, -1, 3, 7}, 3, 11, 1},
		{"all negatives max is least negative window", []int{-4, -2, -3, -1}, 2, -4, 2},
		{"k equals length only one window at 0", []int{1, 2, 3, 4}, 4, 10, 0},
		{"k one max element earliest if tie", []int{3, 5, 4, 5, 2}, 1, 5, 1},
		{"tie for max sum earliest index wins", []int{4, 1, 1, 4, 1}, 2, 5, 0},
		{"k zero invalid", []int{1, 2, 3}, 0, 0, -1},
		{"k negative invalid", []int{1, 2, 3}, -2, 0, -1},
		{"k exceeds length invalid", []int{1, 2, 3}, 5, 0, -1},
		{"empty array invalid", []int{}, 1, 0, -1},
	}
	for _, c := range cases {
		gotSum, gotStart := MaxSumKAndStart(c.a, c.k)
		if gotSum != c.wantSum || gotStart != c.wantStart {
			t.Errorf("%s: MaxSumKAndStart(%v, %d) = (%d, %d), want (%d, %d)",
				c.name, c.a, c.k, gotSum, gotStart, c.wantSum, c.wantStart)
		}
	}
}

func TestMinLenSumAtLeast(t *testing.T) {
	cases := []struct {
		name   string
		a      []int
		target int
		want   int
	}{
		{"classic two element answer", []int{2, 3, 1, 2, 4, 3}, 7, 2},
		{"whole array needed", []int{1, 1, 1, 1}, 4, 4},
		{"prefix exceeds target", []int{1, 2, 3, 4, 5}, 11, 3},
		{"single element reaches target", []int{5}, 5, 1},
		{"single element short", []int{5}, 6, 0},
		{"no subarray reaches target", []int{1, 2, 3, 4, 5}, 100, 0},
		{"empty array", []int{}, 1, 0},
		{"target zero needs nothing", []int{2, 3, 1, 2, 4, 3}, 0, 0},
		{"target negative needs nothing", []int{2, 3, 1, 2, 4, 3}, -5, 0},
		{"single big element beats longer sum", []int{10, 2, 3}, 6, 1},
		{"best window in the middle", []int{1, 4, 4}, 8, 2},
	}
	for _, c := range cases {
		if got := MinLenSumAtLeast(c.a, c.target); got != c.want {
			t.Errorf("%s: MinLenSumAtLeast(%v, %d) = %d, want %d", c.name, c.a, c.target, got, c.want)
		}
	}
}

func intsEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}