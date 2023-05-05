/*
Exercise 11.6
Напишите функцию производительности для сравнения реализации PopCount из раздела
2.6.2 с вашими решениями упражнений 2.4 и 2.5. В какой момент не срабатывает
даже табличное тестирование?
*/

package popcount

import (
	"testing"

	"GolangBook/chapter11/lesson4/ex11.6/popcount"
)

const bin = 0x1234567890ABCDEF

func bench(b *testing.B, f func(uint64) int) {
	for i := 0; i < b.N; i++ {
		f(uint64(bin))
	}
}

func BenchmarkPopCount(b *testing.B) {
	bench(b, popcount.PopCount)
}

func BenchmarkTableLoop(b *testing.B) {
	bench(b, popcount.TableLoop)
}

func BenchmarkPopCountShiftValue(b *testing.B) {
	bench(b, popcount.PopCountShiftValue)
}

func BenchmarkPopCountDiscardBit(b *testing.B) {
	bench(b, popcount.PopCountDiscardBit)
}

/*
go test -bench='.' -benchmem
BenchmarkPopCount-16                    1000000000               0.2177 ns/op          0 B/op          0 allocs/op
BenchmarkTableLoop-16                   420792625                2.709 ns/op           0 B/op          0 allocs/op
BenchmarkPopCountShiftValue-16          74911042                15.98 ns/op            0 B/op          0 allocs/op
BenchmarkPopCountDiscardBit-16          70415923                15.27 ns/op            0 B/op          0 allocs/op
PASS
ok      GolangBook/chapter11/lesson4/ex11.6     4.193s
*/
