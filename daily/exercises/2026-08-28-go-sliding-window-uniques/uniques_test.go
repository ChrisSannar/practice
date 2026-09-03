package uniques

import "testing"

func TestFirstRepeatedIndex(t *testing.T) {
	tests := []struct {
		name string
		a    []int
		want int
	}{
		{"first repeat mid", []int{5, 1, 3, 1, 7}, 3},
		{"no repeats", []int{1, 2, 3}, -1},
		{"empty", []int{}, -1},
		{"single", []int{9}, -1},
		{"adjacent dup", []int{4, 4, 4}, 1},
		{"repeat at end", []int{1, 2, 3, 2}, 3},
		{"zeros and negatives", []int{0, -1, -1, 0}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FirstRepeatedIndex(tt.a); got != tt.want {
				t.Errorf("FirstRepeatedIndex(%v) = %d, want %d", tt.a, got, tt.want)
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
		{"interior repeat", []int{1, 2, 3, 1, 2}, 3},
		{"all same", []int{7, 7, 7}, 1},
		{"empty", []int{}, 0},
		{"single", []int{5}, 1},
		{"all distinct", []int{1, 2, 3, 4, 5}, 5},
		{"repeat forces shrink past several", []int{1, 2, 3, 3, 4, 5}, 3}, // This was originally 4, however that was incorrect.
		{"best window is early", []int{1, 2, 3, 4, 4, 1}, 4},
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
		{"abcabcbb", "abcabcbb", 3},
		{"bbbbb", "bbbbb", 1},
		{"pwwkew", "pwwkew", 3},
		{"empty", "", 0},
		{"single", "a", 1},
		{"all distinct", "abcdef", 6},
		{"repeat at boundary", "abba", 2},
		{"unicode multibyte", "héllo", 3},
		{"unicode all distinct", "café", 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LengthOfLongestSubstring(tt.s); got != tt.want {
				t.Errorf("LengthOfLongestSubstring(%q) = %d, want %d", tt.s, got, tt.want)
			}
		})
	}
}
