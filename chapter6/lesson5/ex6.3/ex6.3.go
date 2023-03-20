/*
Exercise 6.3
(*IntSet).UnionWith вычисляет объединение двух множеств с помощью оператора |, побитового оператора ИЛИ.
Реализуйте методы IntersectWith, DifferenceWith и SymmetricDifference для соответствующих операций над множествами.
(Симметричная разность двух множеств содержит элементы, имеющиеся в одном из множеств, но не в обоих одновременно.)

Проблема:
Не мог понять как получать и работать с битами, какие побитовые операторы использовать. Помогли разобраться статьи:
https://backendinterview.ru/unsorted/bits
https://uproger.com/struktury-dannyh-golang-podrobno/
https://rtfm.co.ua/ru/books-translations/go-s-nulya/go-chast-6-operatory/
*/

package intset

import (
	"bytes"
	"fmt"
)

// IntSet представляет собой множество небольших неотрицательных
// целых чисел. Нулевое значение представляет пустое множество.
type IntSet struct {
	words []uint64
	count int
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
	}
}

// UnionWith делает множество s равным объединению множеств s и t.
// При операции объединения выбираются элементы обоих множеств. Если элемент присутствует в обоих множествах,
// берется только одна его копия
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
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
	c.words = make([]uint64, len(s.words))
	copy(c.words, s.words)
	// or can use
	// for v := range s.words {
	//		c.words[v] = s.words[v]
	//	}
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
		word, bit := x/64, uint(x%64)
		s.words[word] &^= 1 << bit
		s.count--
	}
}

// Len возвращает количество элементов
func (s *IntSet) Len() int {
	return s.count
}

//  Len другая реализация подсчета элементов
// func (s *IntSet) Len() int {
// 	c := 0
// 	for _, v := range s.words {
// 		c += popcount(v)
// 	}
// 	return c
// }
//
// func popcount(x uint64) int {
// 	count := 0
// 	for x != 0 {
// 		count++
// 		x &= x - 1
// 	}
// 	return count
// }

// Has указывает, содержит ли множество неотрицательное значение x
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add добавляет неотрицательное значение x в множество
func (s *IntSet) Add(x int) {
	if !s.Has(x) {
		word, bit := x/64, uint(x%64)
		for word >= len(s.words) {
			s.words = append(s.words, 0)
		}
		s.count++
		s.words[word] |= 1 << bit
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
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}
