package slicesmaps

import (
	"reflect"
	"testing"
)

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

func TestUnique(t *testing.T) {
	cases := []struct {
		s    []int
		want []int
	}{
		{[]int{}, []int{}},
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{1, 1, 2, 1, 3, 2}, []int{1, 2, 3}},
		{[]int{5, 5, 5}, []int{5}},
	}
	for _, c := range cases {
		if got := Unique(c.s); !reflect.DeepEqual(got, c.want) {
			t.Errorf("Unique(%v) = %v, want %v", c.s, got, c.want)
		}
	}
}

func TestIntersection(t *testing.T) {
	cases := []struct {
		a, b, want []int
	}{
		{[]int{1, 2, 3}, []int{2, 3, 4}, []int{2, 3}},
		{[]int{1, 2, 3}, []int{4, 5, 6}, []int{}},
		{[]int{1, 1, 2, 2}, []int{2, 1}, []int{1, 2}}, // dedup, order from a
		{[]int{}, []int{1, 2}, []int{}},
		{[]int{3, 1, 2}, []int{2, 1}, []int{1, 2}}, // order follows a
	}
	for _, c := range cases {
		if got := Intersection(c.a, c.b); !reflect.DeepEqual(got, c.want) {
			t.Errorf("Intersection(%v, %v) = %v, want %v", c.a, c.b, got, c.want)
		}
	}
}

func TestMerge(t *testing.T) {
	cases := []struct {
		name string
		in   []map[string]int
		want map[string]int
	}{
		{"no maps", nil, map[string]int{}},
		{"single map", []map[string]int{{"a": 1, "b": 2}}, map[string]int{"a": 1, "b": 2}},
		{
			"overlapping keys summed",
			[]map[string]int{{"a": 1, "b": 2}, {"b": 3, "c": 4}},
			map[string]int{"a": 1, "b": 5, "c": 4},
		},
		{
			"three maps",
			[]map[string]int{{"x": 1}, {"x": 1}, {"x": 1, "y": 2}},
			map[string]int{"x": 3, "y": 2},
		},
	}
	for _, c := range cases {
		if got := Merge(c.in...); !reflect.DeepEqual(got, c.want) {
			t.Errorf("%s: Merge(%v) = %v, want %v", c.name, c.in, got, c.want)
		}
	}
}
