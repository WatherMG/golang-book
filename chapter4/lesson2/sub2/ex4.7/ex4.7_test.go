package main

import "testing"

var testData = []byte("QWERTYQWERTYQWERTYQWERTYQWERTYQWERTY")

func BenchmarkRev(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reverse(testData)
	}
}

func TestRevUTF8(t *testing.T) {
	s := []byte("Räksmörgås")
	got := string(revUTF8(s))
	want := "sågrömskäR"
	if got != want {
		t.Errorf("got %v, want %v", string(got), want)
	}
}
