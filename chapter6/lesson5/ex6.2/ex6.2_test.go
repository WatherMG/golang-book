package intset

import (
	"testing"
)

const N = 1

func initIntSet() *IntSet {
	s := &IntSet{}
	for x := 0; x < N; x++ {
		s.Add(x)
	}
	return s
}

func TestIntSet_AddAll(t *testing.T) {
	l := []int{105, 106, 107, 108, 109}
	s := initIntSet()
	s.AddAll(101, 102, 103, 104)
	s.AddAll(l...)
	for _, v := range l {
		if got := s.Has(v); !got {
			t.Fatalf("set doesn't contain %d. Set: %s", v, s.String())
		}
	}

}

func TestIntSet_Len(t *testing.T) {
	s := initIntSet()
	if got := s.Len(); got != N {
		t.Fatalf("got %d, want %d", got, N)
	}
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
