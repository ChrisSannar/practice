package basics

import (
	"reflect"
	"testing"
)

func TestMax(t *testing.T) {
	cases := []struct {
		a, b, want int
	}{
		{1, 2, 2},
		{5, 3, 5},
		{-4, -1, -1},
		{7, 7, 7},
		{0, -2, 0},
	}
	for _, c := range cases {
		if got := Max(c.a, c.b); got != c.want {
			t.Errorf("Max(%d, %d) = %d, want %d", c.a, c.b, got, c.want)
		}
	}
}

func TestSum(t *testing.T) {
	cases := []struct {
		nums []int
		want int
	}{
		{[]int{}, 0},
		{[]int{5}, 5},
		{[]int{1, 2, 3, 4}, 10},
		{[]int{-2, 2, -5}, -5},
	}
	for _, c := range cases {
		if got := Sum(c.nums); got != c.want {
			t.Errorf("Sum(%v) = %d, want %d", c.nums, got, c.want)
		}
	}
}

func TestReverse(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"", ""},
		{"a", "a"},
		{"hello", "olleh"},
		{"Go!", "!oG"},
		{"héllo", "olléh"}, // multi-byte rune must stay intact
	}
	for _, c := range cases {
		if got := Reverse(c.in); got != c.want {
			t.Errorf("Reverse(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}

func TestCountVowels(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"", 0},
		{"xyz", 0},
		{"hello", 2},
		{"AEIOU", 5},
		{"Go Gophers", 3},
	}
	for _, c := range cases {
		if got := CountVowels(c.in); got != c.want {
			t.Errorf("CountVowels(%q) = %d, want %d", c.in, got, c.want)
		}
	}
}

func TestWordCount(t *testing.T) {
	cases := []struct {
		in   string
		want map[string]int
	}{
		{"", map[string]int{}},
		{"   ", map[string]int{}},
		{"a b a", map[string]int{"a": 2, "b": 1}},
		{"the cat the dog the", map[string]int{"the": 3, "cat": 1, "dog": 1}},
	}
	for _, c := range cases {
		if got := WordCount(c.in); !reflect.DeepEqual(got, c.want) {
			t.Errorf("WordCount(%q) = %v, want %v", c.in, got, c.want)
		}
	}
}

func TestFizzBuzz(t *testing.T) {
	cases := []struct {
		n    int
		want []string
	}{
		{0, []string{}},
		{1, []string{"1"}},
		{5, []string{"1", "2", "Fizz", "4", "Buzz"}},
		{15, []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}},
	}
	for _, c := range cases {
		if got := FizzBuzz(c.n); !reflect.DeepEqual(got, c.want) {
			t.Errorf("FizzBuzz(%d) = %v, want %v", c.n, got, c.want)
		}
	}
}
