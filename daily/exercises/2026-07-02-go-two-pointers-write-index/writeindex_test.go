package writeindex

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	cases := []struct {
		name    string
		s       []int
		wantK   int
		wantTop []int // s[:k] after the call
	}{
		{"one dup", []int{1, 1, 2}, 2, []int{1, 2}},
		{"many dups", []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5, []int{0, 1, 2, 3, 4}},
		{"already unique", []int{1, 2, 3}, 3, []int{1, 2, 3}},
		{"empty", []int{}, 0, []int{}},
		{"single", []int{5}, 1, []int{5}},
		{"all same", []int{7, 7, 7}, 1, []int{7}},
	}
	for _, c := range cases {
		k := RemoveDuplicates(c.s)
		if k != c.wantK {
			t.Errorf("%s: RemoveDuplicates(...) = %d, want %d", c.name, k, c.wantK)
			continue
		}
		if got := c.s[:k]; !reflect.DeepEqual(got, c.wantTop) {
			t.Errorf("%s: after RemoveDuplicates, s[:%d] = %v, want %v", c.name, k, got, c.wantTop)
		}
	}
}

func TestMoveZeroes(t *testing.T) {
	cases := []struct {
		name string
		s    []int
		want []int
	}{
		{"interleaved", []int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
		{"leading zeroes", []int{0, 0, 1}, []int{1, 0, 0}},
		{"no zeroes", []int{1, 2, 3}, []int{1, 2, 3}},
		{"all zeroes", []int{0, 0, 0}, []int{0, 0, 0}},
		{"empty", []int{}, []int{}},
		{"trailing zero", []int{1, 0}, []int{1, 0}},
		{"several", []int{4, 0, 5, 0, 0, 6}, []int{4, 5, 6, 0, 0, 0}},
	}
	for _, c := range cases {
		MoveZeroes(c.s)
		if !reflect.DeepEqual(c.s, c.want) {
			t.Errorf("%s: after MoveZeroes, slice = %v, want %v", c.name, c.s, c.want)
		}
	}
}

func TestRemoveElement(t *testing.T) {
	cases := []struct {
		name    string
		s       []int
		val     int
		wantK   int
		wantTop []int // s[:k] after the call
	}{
		{"remove some", []int{3, 2, 2, 3}, 3, 2, []int{2, 2}},
		{"remove from middle", []int{0, 1, 2, 2, 3, 0, 4, 2}, 2, 5, []int{0, 1, 3, 0, 4}},
		{"none match", []int{1, 2, 3}, 9, 3, []int{1, 2, 3}},
		{"all match", []int{4, 4, 4}, 4, 0, []int{}},
		{"empty", []int{}, 1, 0, []int{}},
		{"single match", []int{7}, 7, 0, []int{}},
	}
	for _, c := range cases {
		k := RemoveElement(c.s, c.val)
		if k != c.wantK {
			t.Errorf("%s: RemoveElement(..., %d) = %d, want %d", c.name, c.val, k, c.wantK)
			continue
		}
		if got := c.s[:k]; !reflect.DeepEqual(got, c.wantTop) {
			t.Errorf("%s: after RemoveElement(..., %d), s[:%d] = %v, want %v", c.name, c.val, k, got, c.wantTop)
		}
	}
}
