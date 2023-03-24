package main

import "testing"

func TestWordCounter(t *testing.T) {
	var tcs = []struct {
		input  []byte
		except int
	}{
		{[]byte(""), 0},
		{[]byte("one"), 1},
		{[]byte("one two"), 2},
		{[]byte("one two\nthree four"), 4},
		{[]byte("one two\nthree\nfour\n"), 4},
		{[]byte("one\ntwo\nthree\nfour\nfive"), 5},
	}
	for _, tc := range tcs {
		var c WordCounter
		c.Write(tc.input)
		if int(c) != tc.except {
			t.Errorf("Words of %s, except %d, got %d", tc.input, tc.except, c)
		}
	}
}

func TestLineCounter(t *testing.T) {
	var tcs = []struct {
		input  []byte
		except int
	}{
		{[]byte(""), 0},
		{[]byte("one"), 1},
		{[]byte("one\n  two"), 2},
		{[]byte("one \ntwo\n three \nfour"), 4},
		{[]byte("one \ntwo\n\nfour\n"), 4},
	}
	for _, tc := range tcs {
		var l LineCounter
		l.Write(tc.input)
		if int(l) != tc.except {
			t.Errorf("Words of %s, except %d, got %d", tc.input, tc.except, l)
		}
	}
}
