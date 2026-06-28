package windows

import (
	"reflect"
	"testing"
)

func TestTake(t *testing.T) {
	cases := []struct {
		s    []int
		n    int
		want []int
	}{
		{[]int{1, 2, 3, 4}, 2, []int{1, 2}},
		{[]int{1, 2, 3}, 0, []int{}},
		{[]int{1, 2, 3}, -1, []int{}},
		{[]int{1, 2}, 5, []int{1, 2}}, // n bigger than len -> all
		{[]int{}, 3, []int{}},
	}
	for _, c := range cases {
		if got := Take(c.s, c.n); !reflect.DeepEqual(got, c.want) {
			t.Errorf("Take(%v, %d) = %v, want %v", c.s, c.n, got, c.want)
		}
	}
}

func TestDrop(t *testing.T) {
	cases := []struct {
		s    []int
		n    int
		want []int
	}{
		{[]int{1, 2, 3, 4}, 2, []int{3, 4}},
		{[]int{1, 2, 3}, 0, []int{1, 2, 3}},
		{[]int{1, 2, 3}, -1, []int{1, 2, 3}},
		{[]int{1, 2}, 5, []int{}}, // n bigger than len -> empty
		{[]int{}, 3, []int{}},
	}
	for _, c := range cases {
		if got := Drop(c.s, c.n); !reflect.DeepEqual(got, c.want) {
			t.Errorf("Drop(%v, %d) = %v, want %v", c.s, c.n, got, c.want)
		}
	}
}

func TestChunk(t *testing.T) {
	cases := []struct {
		name string
		s    []int
		size int
		want [][]int
	}{
		{"empty", []int{}, 3, [][]int{}},
		{"exact multiple", []int{1, 2, 3, 4}, 2, [][]int{{1, 2}, {3, 4}}},
		{"ragged last chunk", []int{1, 2, 3, 4, 5}, 2, [][]int{{1, 2}, {3, 4}, {5}}},
		{"size bigger than len", []int{1, 2}, 5, [][]int{{1, 2}}},
		{"size of one", []int{1, 2, 3}, 1, [][]int{{1}, {2}, {3}}},
		{"zero size yields nothing", []int{1, 2, 3}, 0, [][]int{}},
		{"negative size yields nothing", []int{1, 2, 3}, -2, [][]int{}},
	}
	for _, c := range cases {
		if got := Chunk(c.s, c.size); !reflect.DeepEqual(got, c.want) {
			t.Errorf("%s: Chunk(%v, %d) = %v, want %v", c.name, c.s, c.size, got, c.want)
		}
	}
}
