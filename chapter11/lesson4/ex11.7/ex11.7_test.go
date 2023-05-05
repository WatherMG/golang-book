/*
Exercise 11.7
Напишите функции производительности для Add, UnionWith и других
методов *IntSet (раздел 6.5) с использованием больших псевдослучайных входных
данных. Насколько быстрыми вы сможете сделать эти методы? Как влияет на
производительность выбор размера слова? Насколько быстро работает IntSet по
сравнению с реализацией множества на основе отображения?
*/

package intset

import (
	"math/rand"
	"testing"
	"time"

	"GolangBook/chapter11/lesson4/ex11.7/intset"
)

var (
	s1 []int
	s2 []int
)

const (
	n     = 100000
	scale = 100
)

func init() {
	seed := time.Now().UTC().UnixNano()
	rand.New(rand.NewSource(seed))
	s1 = randInt(n)
	s2 = randInt(n)
}

func randInt(n int) []int {
	ints := make([]int, n)
	for i := 0; i < n; i++ {
		ints[i] = rand.Intn(scale * n)
	}
	return ints
}

func BenchmarkIntSetAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := &intset.IntSet{}
		for _, v := range s1 {
			s.Add(v)
		}
	}
}

func BenchmarkIntSetHas(b *testing.B) {
	s := &intset.IntSet{}
	for _, v := range s1 {
		s.Add(v)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, v := range s1 {
			s.Has(v)
		}
	}
}

func BenchmarkIntSetUnionWith(b *testing.B) {
	is1 := &intset.IntSet{}
	for _, v := range s1 {
		is1.Add(v)
	}
	is2 := &intset.IntSet{}
	for _, v := range s2 {
		is2.Add(v)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		is1.UnionWith(is2)
	}
}

func BenchmarkMapAdd(b *testing.B) {
	s := make(map[int]bool)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, k := range s1 {
			s[k] = true
		}
	}
}

func BenchmarkMapHas(b *testing.B) {
	s := make(map[int]bool)
	for _, k := range s1 {
		s[k] = true
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, k := range s1 {
			_ = s[k]
		}
	}
}

func BenchmarkMapUnionWith(b *testing.B) {
	ms1 := make(map[int]bool)
	for _, k := range s1 {
		ms1[k] = true
	}
	ms2 := make(map[int]bool)
	for _, k := range s2 {
		ms2[k] = true
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for k := range ms2 {
			ms1[k] = true
		}
	}
}
