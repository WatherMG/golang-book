package intset

import "testing"

type interval struct {
	start, end int
}

func initIntSet(i interval) *IntSet {
	s := &IntSet{}
	for x := i.start; x < i.end; x++ {
		s.Add(x)

	}
	return s
}

func initMapIntSet(i interval) map[int]bool {
	s := make(map[int]bool)
	for x := i.start; x < i.end; x++ {
		s[x] = true
	}
	return s
}

func equal(s1 *IntSet, s2 map[int]bool) bool {
	for k := range s2 {
		if !s1.Has(k) {
			return false
		}
	}
	return true
}

func TestIntSet(t *testing.T) {
	tests := []struct {
		seti interval
		mapi interval
		want bool
	}{
		{interval{0, 50}, interval{0, 50}, true},
		{interval{0, 100}, interval{0, 100}, true},
		{interval{0, 49}, interval{0, 50}, false},
	}
	for i, tt := range tests {
		s1 := initIntSet(tt.seti)
		s2 := initMapIntSet(tt.mapi)
		if got := equal(s1, s2); got != tt.want {
			t.Errorf("TestIntSet: %d. got %v, want %v", i, got, tt.want)
		}
	}
}

func TestIntSetUnionWith(t *testing.T) {
	tests := []struct {
		seti1, seti2 interval
		mapi1, mapi2 interval
		want         bool
	}{
		{interval{0, 50}, interval{50, 100},
			interval{0, 50}, interval{50, 100}, true},

		{interval{0, 50}, interval{50, 99},
			interval{0, 50}, interval{0, 100}, false},
	}
	for i, test := range tests {
		s1, s2 := initIntSet(test.seti1), initIntSet(test.seti2)
		s1.UnionWith(s2)

		ms1, ms2 := initMapIntSet(test.mapi1), initMapIntSet(test.mapi2)
		for k := range ms2 {
			ms1[k] = true
		}
		if got := equal(s1, ms1); got != test.want {
			t.Errorf("TestIntSetUnionWith: %d. got %v, want %v", i, got, test.want)
		}
	}
}
