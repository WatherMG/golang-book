package main

import "testing"

func Test_countDifferenceBits(t *testing.T) {
	tests := []struct {
		a, b [32]byte
		want int
	}{
		{[32]byte{0}, [32]byte{6}, 2},
		{[32]byte{1, 2, 3}, [32]byte{4, 5, 6}, 7},
		{[32]byte{0: 0x81, 0x82}, [32]byte{0: 0x82, 0x81}, 4},
		{[32]byte{0: 0x81}, [32]byte{0: 0x82}, 2},
		{[32]byte{0: 0x81}, [32]byte{0: 0x82, 0xcd, 0xdd}, 13},
	}
	for _, test := range tests {
		got := countDifferenceBits(test.a, test.b)
		if got != test.want {
			t.Errorf("bitDiff(%v, %v), got %d, want %d",
				test.a, test.b, got, test.want)
		}
	}
}
