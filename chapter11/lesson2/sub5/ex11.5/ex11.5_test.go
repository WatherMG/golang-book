package main

import (
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	tests := []struct {
		s    string
		sep  string
		want int
	}{
		{"a b c", " ", 3},
		{"a.b.c", ".", 3},
		{"a b c d e", " ", 5},
		{"a b cde", " ", 3},
		{"a-b-cd-ef-g", "-", 5},
		{"abc", "", 3},
		{"a_b_c", "_", 3},
	}
	for _, test := range tests {
		words := strings.Split(test.s, test.sep)
		if got := len(words); got != test.want {
			t.Errorf("Split(%q, %q) возвращает %d слов, а требуется %d", test.s, test.sep, got, test.want)
		}
	}
}
