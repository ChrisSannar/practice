package main

import (
	"testing"
)

func TestEnglishWordToInteger(tester *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"One Hundred Twenty Three", 123},
		{"Twelve Thousand Three Hundred Forty Five", 12345},
		{"One Million Two Hundred Thirty Four Thousand Five Hundred Sixty Seven", 1234567},
		{"Zero", 0},
	}

	for _, c := range cases {
		got := EnglishWordToInt(c.in)
		if got != c.want {
			tester.Errorf("EnglishWordToInt(%q) == %d, want %d", c.in, got, c.want)
		}
	}
}

func TestIntToEnglishWord(tester *testing.T) {
	cases := []struct {
		in   int
		want string
	}{
		{123, "One Hundred Twenty Three"},
		{12345, "Twelve Thousand Three Hundred Forty Five"},
		{1234567, "One Million Two Hundred Thirty Four Thousand Five Hundred Sixty Seven"},
		{0, "Zero"},
	}

	for _, c := range cases {
		got := IntToEnglishWord(c.in)
		if got != c.want {
			tester.Errorf("EnglishWordToInt(%d) == %q, want %q", c.in, got, c.want)
		}
	}
}
