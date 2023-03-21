package intset

import (
	"math/rand"
	"testing"
	"time"
)

const (
	MIN = iota << 4
	N
	MAX
)

func initIntSet(n int) *IntSet {
	s := &IntSet{}
	for x := 0; x < n; x++ {
		s.Add(x)
	}
	return s
}

func TestIntSet_Elems(t *testing.T) {
	s := &IntSet{}
	l := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s.AddAll(l...)
	for i, v := range s.Elems() {
		match := l[i] == v
		if !match {
			t.Logf("Elems: value error: s[%d]: %d != l[%d]: %d", i, v, i, l[i])
			t.Fail()
		}
	}
}

func TestIntSet_SymmetricDifference(t *testing.T) {
	s := &IntSet{}
	k := &IntSet{}

	rand.NewSource(time.Now().UnixNano())

	for i := 0; i < N; i++ {
		s.Add(rand.Intn(MAX-MIN) + MIN)
		k.Add(rand.Intn(MAX-MIN) + MIN)
	}
	c := s.Copy()

	// t.Logf("s = %s", c)
	// t.Logf("t = %s", k)
	// t.Logf("s.count=%d", c.count)
	c.SymmetricDifference(k)
	// t.Logf("res: %s", c)
	// t.Logf("SymmetricDifference: s.count=%d", c.count)

	// Симметричная разность двух множеств содержит элементы, которые есть в одном множестве, но не в обоих одновременно.
	for _, elem := range c.Elems() {
		inC := s.Has(elem)
		inK := k.Has(elem)
		if inC == inK {
			t.Errorf("SymmetricDifference: %d contains in s and t", elem)
		}
	}
}

func TestIntSet_DifferenceWith(t *testing.T) {
	s := &IntSet{}
	k := &IntSet{}

	rand.NewSource(time.Now().UnixNano())

	for i := 0; i < N; i++ {
		s.Add(rand.Intn(MAX-MIN) + MIN)
		k.Add(rand.Intn(MAX-MIN) + MIN)
	}

	c := s.Copy()

	// t.Log(s)
	// t.Log(k)
	// t.Logf("s.count=%d", s.count)
	s.DifferenceWith(k)
	// t.Log(s)
	// t.Logf("DifferenceWith: s.count=%d", s.count)

	for _, i := range s.Elems() {
		if c.Has(i) && s.Has(i) && k.Has(i) {
			t.Errorf("DifferenceWith: difference is incorect: %s", s)
		}
		if c.Has(i) && !s.Has(i) && !k.Has(i) {
			t.Errorf("DifferenceWith: difference is incorect: %s", s)
		}
	}
}

func TestIntSet_IntersectWith(t *testing.T) {
	s := &IntSet{}
	k := &IntSet{}

	rand.NewSource(time.Now().UnixNano())

	for i := 0; i < N; i++ {
		s.Add(rand.Intn(MAX-MIN) + MIN)
		k.Add(rand.Intn(MAX-MIN) + MIN)
	}
	// t.Log(s)
	// t.Log(k)

	// t.Logf("s.count=%d", s.count)
	s.IntersectWith(k)
	// t.Log(s)
	// t.Logf("IntersectWith: s.count=%d", s.count)

	// Проверка, что результат содержит только элементы, присутствующе в обоих множествах
	for _, i := range s.Elems() {
		if s.Has(i) != (k.Has(i) && s.Has(i)) {
			t.Fail()
		}
	}
}

func TestIntSet_UnionWith(t *testing.T) {
	s := &IntSet{}
	k := &IntSet{}

	rand.NewSource(time.Now().UnixNano())

	for i := 0; i < N; i++ {
		s.Add(rand.Intn(MAX-MIN) + MIN)
		k.Add(rand.Intn(MAX-MIN) + MIN)
	}
	// t.Log(s)
	// t.Log(k)

	// t.Logf("s.count=%d", s.count)
	s.UnionWith(k)
	// t.Log(s)
	// t.Logf("UnionWith: s.count=%d", s.count)
	for _, i := range s.Elems() {
		if !s.Has(i) && k.Has(i) {
			t.Fail()
		}
	}
}

func TestIntSet_AddAll(t *testing.T) {
	l := []int{105, 106, 107, 108, 109}
	s := initIntSet(N)
	s.AddAll(101, 102, 103, 104)
	s.AddAll(l...)
	for _, v := range l { // or check []int{101, 102, 103, 104}
		if got := s.Has(v); !got {
			t.Fatalf("set doesn't contain %d. Set: %s", v, s.String())
		}
	}
}

func TestIntSet_Len(t *testing.T) {
	s := initIntSet(N)
	if got := s.Len(); got != N {
		t.Fatalf("got %d, want %d", got, N)
	}
}

func TestIntSet_Remove(t *testing.T) {
	s := initIntSet(N)
	n := 1
	s.Remove(n)
	if got := s.Has(n); got {
		t.Fatalf("didn't remove %d", n)
	}
}

func TestIntSet_Clear(t *testing.T) {
	s := initIntSet(N)
	s.Clear()
	if got := s.Len(); got != 0 {
		t.Fatalf("didn't clear set, len = %d", s.Len())
	}
}

func TestIntSet_Copy(t *testing.T) {
	s := initIntSet(N)
	c := s.Copy()
	for v := range s.words {
		if c.words[v] != s.words[v] {
			t.Fatalf("%s is not %s", s.String(), c.String())
		}
	}
}
