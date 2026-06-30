package twopointers

import (
	"reflect"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	cases := []struct {
		name string
		s    []int
		want bool
	}{
		{"odd palindrome", []int{1, 2, 3, 2, 1}, true},
		{"even palindrome", []int{1, 2, 2, 1}, true},
		{"not a palindrome", []int{1, 2, 3}, false},
		{"near miss at ends", []int{1, 2, 3, 4}, false},
		{"empty", []int{}, true},
		{"single", []int{5}, true},
		{"two equal", []int{7, 7}, true},
		{"two different", []int{1, 2}, false},
	}
	for _, c := range cases {
		if got := IsPalindrome(c.s); got != c.want {
			t.Errorf("%s: IsPalindrome(%v) = %v, want %v", c.name, c.s, got, c.want)
		}
	}
}

func TestReverseInPlace(t *testing.T) {
	cases := []struct {
		name string
		s    []int
		want []int
	}{
		{"odd length", []int{1, 2, 3}, []int{3, 2, 1}},
		{"even length", []int{1, 2, 3, 4}, []int{4, 3, 2, 1}},
		{"single", []int{5}, []int{5}},
		{"empty", []int{}, []int{}},
		{"two", []int{1, 2}, []int{2, 1}},
	}
	for _, c := range cases {
		ReverseInPlace(c.s)
		if !reflect.DeepEqual(c.s, c.want) {
			t.Errorf("%s: after ReverseInPlace, slice = %v, want %v", c.name, c.s, c.want)
		}
	}
}
