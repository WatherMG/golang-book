/*
Exercise 9.2
Перепишите пример PopCount из раздела 2.6.2 так, чтобы он
инициализировал таблицу поиска с использованием sync.Once при первом к ней
обращении. (В реальности стоимость синхронизации для таких малых и
высокооптимизированных функций, как PopCount, является чрезмерно высокой.)
*/

package popcount

import "sync"

var initTableOnce sync.Once

// pc[i] - количество единичных битов в i.
var pc [256]byte

func loadTable() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount возвращает степень заполнения (количество установленных битов) значения x.
func PopCount(x uint64) int {
	initTableOnce.Do(loadTable)
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}
