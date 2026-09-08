package shrink

import "testing"

func TestShrinkPastDuplicate(t *testing.T) {
	tests := []struct {
		name        string
		a           []int
		left, right int
		want        int
	}{
		{"duplicate several steps in", []int{1, 2, 5, 9, 5}, 0, 4, 3},
		{"duplicate right at left", []int{5, 1, 2, 9, 5}, 0, 4, 1},
		{"no duplicate, left stays put", []int{1, 2, 3, 4, 5}, 0, 4, 0},
		{"no duplicate, nonzero left stays put", []int{9, 1, 2, 3, 4, 5}, 1, 5, 1},
		{"duplicate is the element right before right", []int{1, 9, 2, 3, 3}, 0, 4, 4},
		{"single-element window before right, matches", []int{3, 3}, 0, 1, 1},
		{"single-element window before right, no match", []int{3, 4}, 0, 1, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ShrinkPastDuplicate(tt.a, tt.left, tt.right); got != tt.want {
				t.Errorf("ShrinkPastDuplicate(%v, %d, %d) = %d, want %d", tt.a, tt.left, tt.right, got, tt.want)
			}
		})
	}
}

func TestLongestUniqueLen(t *testing.T) {
	tests := []struct {
		name string
		a    []int
		want int
	}{
		{"empty", []int{}, 0},
		{"single", []int{5}, 1},
		{"all same", []int{7, 7, 7}, 1},
		{"all distinct", []int{1, 2, 3, 4}, 4},
		{"interior repeat, adjacent-ish", []int{1, 2, 3, 1, 2}, 3},
		{"non-adjacent repeat forces multi-step shrink", []int{0, 1, 2, 0, 3}, 4},
		{"repeat forces shrink past several", []int{1, 2, 3, 4, 2, 5}, 4},
		{"best window is early", []int{1, 2, 3, 1, 1}, 3},
		{"best window is late", []int{1, 1, 2, 3, 4}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LongestUniqueLen(tt.a); got != tt.want {
				t.Errorf("LongestUniqueLen(%v) = %d, want %d", tt.a, got, tt.want)
			}
		})
	}
}

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"empty", "", 0},
		{"single", "a", 1},
		{"all same", "bbbbb", 1},
		{"classic LC3 case", "abcabcbb", 3},
		{"non-adjacent repeat, the real trap", "dvdf", 3},
		{"pwwkew", "pwwkew", 3},
		{"longer non-adjacent shrink", "tmmzuxt", 5},
		{"all distinct", "abcdef", 6},
		{"unicode multibyte, all distinct", "日本語abc", 6},
		{"unicode multibyte, forces shrink", "日本語日本", 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LengthOfLongestSubstring(tt.s); got != tt.want {
				t.Errorf("LengthOfLongestSubstring(%q) = %d, want %d", tt.s, got, tt.want)
			}
		})
	}
}
