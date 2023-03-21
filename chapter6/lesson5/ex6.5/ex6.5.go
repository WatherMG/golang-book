/*
Exercise 6.5
Типом каждого слова, используемого в IntSet, является uint64, но 64-разрядная арифметика может быть неэффективной на
32-разрядных платформах. Измените программу так, чтобы она использовала тип uint, который представляет собой наиболее
эффективный беззнаковый целочисленный тип для данной платформы. Вместо деления на 64 определите константу, в которой
хранится эффективный размер uint в битах, 32 или 64. Для этого можно воспользоваться, возможно, слишком умным выражением
32<<(^uint(0)>>63).
*/

package intset

import (
	"bytes"
	"fmt"
)

// PLATFORM определяет разрядность платформы 32|64 в зависимости от размера uint на текущей платформе
const PLATFORM = 32 << (^uint(0) >> 63)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount Пример из раздела 2.6.2 (../chapter2/lesson6/sub2/popcount.go)
func PopCount(x uint) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// IntSet представляет собой множество небольших неотрицательных
// целых чисел. Нулевое значение представляет пустое множество.
type IntSet struct {
	words []uint
	count int
}

func (s *IntSet) Elems() (result []int) {
	result = make([]int, 0, s.Len())
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < PLATFORM; j++ {
			if word&(1<<uint(j)) != 0 {
				result = append(result, PLATFORM*i+j)
			}
		}
	}
	return result
}

// SymmetricDifference делает множество s равным симметричной разнице множеств s и t
// Симметричная разность двух множеств содержит элементы, которые есть в одном множестве, но не в обоих одновременно.
func (s *IntSet) SymmetricDifference(t *IntSet) {
	c := s.Copy()
	s.UnionWith(t)
	c.IntersectWith(t)
	s.DifferenceWith(c)
}

// // SymmetricDifference другая реализация
// func (s *IntSet) SymmetricDifference(t *IntSet) {
// 	for i, tword := range t.words {
// 		if i < len(s.words) {
// 			s.words[i] ^= tword
// 		} else {
// 			s.words = append(s.words, tword)
// 		}
// 		s.count = PopCount(s.words[i])
// 	}
// }

// DifferenceWith делает множество s равным разнице множеств s и t
// Разность множеств — это все элементы, которые содержатся в одном множестве (где вызывается метод),
// но не содержатся в другом (куда передается аргументом).
func (s *IntSet) DifferenceWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &^= tword
		} else {
			s.words = append(s.words, tword)
		}
		s.count = PopCount(s.words[i])
	}
}

// IntersectWith делает множество s равным пересечению множеств s и t.
// Пересечение множеств содержит только те элементы, которые есть в обоих множествах.
func (s *IntSet) IntersectWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= tword
		} else {
			s.words = append(s.words, tword)
		}
		s.count = PopCount(s.words[i])
	}
}

// UnionWith делает множество s равным объединению множеств s и t.
// При операции объединения выбираются элементы обоих множеств. Если элемент присутствует в обоих множествах,
// берется только одна его копия.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
		s.count = PopCount(s.words[i])
	}
}

// AddAll добавляет список значений в множество
func (s *IntSet) AddAll(list ...int) {
	for _, v := range list {
		s.Add(v)
	}
}

// Copy возвращает копию множества
func (s *IntSet) Copy() *IntSet {
	c := &IntSet{}
	c.words = make([]uint, len(s.words))
	copy(c.words, s.words)
	c.count = s.count
	return c
}

// Clear удаляет все элементы из множества
func (s *IntSet) Clear() {
	if s.Len() > 0 {
		s.count = 0
		s.words = nil
	}
}

// Remove удаляет x из множества
func (s *IntSet) Remove(x int) {
	if s.Len() != 0 && s.Has(x) {
		word, bit := x/PLATFORM, uint(x%PLATFORM)
		s.words[word] &^= 1 << bit
		s.count--
	}
}

// Len возвращает количество элементов
func (s *IntSet) Len() int {
	return s.count
}

// // Len другая реализация подсчета элементов
// func (s *IntSet) Len() int {
// 	c := 0
// 	for _, v := range s.words {
// 		c += PopCount(v)
// 	}
// 	return c
// }

// Has указывает, содержит ли множество неотрицательное значение x
func (s *IntSet) Has(x int) bool {
	word, bit := x/PLATFORM, uint(x%PLATFORM)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add добавляет неотрицательное значение x в множество
func (s *IntSet) Add(x int) {
	if !s.Has(x) {
		word, bit := x/PLATFORM, uint(x%PLATFORM)
		for word >= len(s.words) {
			s.words = append(s.words, 0)
		}
		s.words[word] |= 1 << bit
		s.count++
	}
}

// String возвращает множество как строку вида "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < PLATFORM; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", PLATFORM*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}
