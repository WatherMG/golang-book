package intset

import (
	"testing"
)

const N = 100

func initIntSet() *IntSet {
	s := &IntSet{}
	for x := 0; x < N; x++ {
		s.Add(x)
	}
	return s
}

func TestIntSet_Len(t *testing.T) {
	s := initIntSet()
	if got := s.Len(); got != N {
		t.Fatalf("got %d, want %d", got, N)
	}
	t.Logf("len=%d", s.Len())
}

func TestIntSet_Remove(t *testing.T) {
	s := initIntSet()
	n := 1
	s.Remove(n)
	if got := s.Has(n); got {
		t.Fatalf("didn't remove %d", n)
	}
}

func TestIntSet_Clear(t *testing.T) {
	s := initIntSet()
	s.Clear()
	if got := s.Len(); got != 0 {
		t.Fatalf("didn't clear set, len = %d", s.Len())
	}
}

func TestIntSet_Copy(t *testing.T) {
	s := initIntSet()
	c := s.Copy()
	for v := range s.words {
		if c.words[v] != s.words[v] {
			t.Fatalf("%s is not %s", s.String(), c.String())
		}
	}
}
